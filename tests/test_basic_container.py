"""Tests for basic container."""
import os
from unittest import TestCase
import docker


def run_container(
    docker_env: docker.DockerClient, current_directory: str
) -> docker.DockerClient.containers:
    """Create container function."""
    try:
        docker_env.images.get("qatcomputer")
    except docker.errors.ImageNotFound:
        build_image(docker_env=docker_env, current_directory=current_directory)

    return docker_env.containers.run(
        image="qatcomputer",
        name="qatcompter",
        auto_remove=True,
        cgroupns="host",
        detach=True,
        privileged=True,
        volumes=[
            "/sys/fs/cgroup:/sys/fs/cgroup:rw",
            os.path.join(
                current_directory,
                "../qat_computer/resources/conf:/etc/qat-computer/conf:rw",
            ),
            os.path.join(
                current_directory,
                "../qat_computer/resources/compute:/etc/qat-computer/compute:rw",
            ),
        ],
    )


def build_image(docker_env: docker.DockerClient, current_directory: str):
    """Build the image function."""
    docker_env.images.build(
        path=os.path.join(current_directory, "../"),
        dockerfile=os.path.join(current_directory, "../docker/Dockerfile"),
        tag="qatcomputer",
    )


class TestBasicContainer(TestCase):
    """TestBasicContainer."""

    def setUp(self) -> None:
        """SetUp container object."""
        self.current_directory = os.path.dirname(os.path.abspath(__file__))
        self.build_img = False
        self.binary = "qat-computer"
        self.noflagcmd = ""
        self.client = docker.from_env()
        self.container = run_container(
            docker_env=self.client, current_directory=self.current_directory
        )

    def test_version_endpoint(self):
        """Test version endpoint."""
        self.noflagcmd = "-version"
        exit_code, output = self.container.exec_run(
            cmd=f"{self.binary} {self.noflagcmd}", tty=True
        )

        self.assertEqual(exit_code, 0)
        with open(
            os.path.join(self.current_directory, "../VERSION.txt"),
            "r",
            encoding="utf-8",
        ) as vers_file:
            self.assertTrue(bytes(vers_file.read(), "utf-8") in output)

#    def test_qiskit_version_endpoint(self):
#        """Test qiskit version endpoint."""
#        self.noflagcmd = "-qiskit-version"
#        exit_code, output = self.container.exec_run(
#            cmd=f"{self.binary} {self.noflagcmd}", tty=True
#        )
#
#        self.assertEqual(exit_code, 0)
#        self.assertTrue(b"qiskit" in output)

    def test_os_version_endpoint(self):
        """Test os version endpoint."""
        self.noflagcmd = "-os-version"
        exit_code, output = self.container.exec_run(
            cmd=f"{self.binary} {self.noflagcmd}", tty=True
        )

        self.assertEqual(exit_code, 0)
        self.assertTrue(b"Ubuntu" in output)

    def test_conf_endpoint(self):
        """Test conf endpoint."""
        self.noflagcmd = "-show-config -conf=conf_docker.yaml"
        exit_code, output = self.container.exec_run(
            cmd=f"{self.binary} {self.noflagcmd}", tty=True
        )

        self.assertEqual(exit_code, 0)
        self.assertTrue(
            b'"ConfPath":"/etc/qat-computer/conf/conf_docker.yaml"' in output
        )

    def tearDown(self) -> None:
        """TearDown container object."""
        self.container.remove(force=True)
