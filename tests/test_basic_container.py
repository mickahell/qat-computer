"""Tests for basic container."""
import os
import docker
from unittest import TestCase


def run_container(
    docker_env: docker.DockerClient, current_directory: str
) -> docker.DockerClient.containers:
    """Create container function."""
    return docker_env.containers.run(
        image="qatcomputer",
        name="qatcompter",
        auto_remove=True,
        cgroupns="host",
        detach=True,
        privileged=True,
        volumes=[
            "/sys/fs/cgroup:/sys/fs/cgroup:rw",
            f"{current_directory}/../qat_computer/resources/conf:/etc/qat-computer/conf",
            f"{current_directory}/../qat_computer/resources/compute:/etc/qat-computer/compute",
        ],
    )


def build_image(docker_env: docker.DockerClient, current_directory: str):
    """Build the image function."""
    img, _ = docker_env.images.build(
        path=f"{current_directory}/../",
        dockerfile=f"{current_directory}/../docker/Dockerfile",
        tag="qatcomputer",
    )


class TestBasicContainer(TestCase):
    """TestBasicContainer."""

    def setUp(self) -> None:
        """SetUp container object."""
        self.current_directory = os.path.dirname(os.path.abspath(__file__))
        self.build_img = False
        self.BINARY = "qat-computer"
        self.NOFLAGCMD = ""
        self.client = docker.from_env()
        try:
            self.client.images.get("qatcomputer")
        except:
            build_image(
                docker_env=self.client, current_directory=self.current_directory
            )
        self.container = run_container(
            docker_env=self.client, current_directory=self.current_directory
        )

    def test_version_endpoint(self):
        """Test version endpoint."""
        self.NOFLAGCMD = "-version"
        exit_code, output = self.container.exec_run(
            cmd=f"{self.BINARY} {self.NOFLAGCMD}", tty=True
        )

        self.assertEqual(exit_code, 0)
        with open(
            os.path.join(self.current_directory, "../VERSION.txt"),
            "r",
            encoding="utf-8",
        ) as vers_file:
            self.assertTrue(bytes(vers_file.read(), "utf-8") in output)

    def test_qiskit_version_endpoint(self):
        """Test qiskit version endpoint."""
        self.NOFLAGCMD = "-qiskit-version"
        exit_code, output = self.container.exec_run(
            cmd=f"{self.BINARY} {self.NOFLAGCMD}", tty=True
        )

        self.assertEqual(exit_code, 0)
        self.assertTrue(b"qiskit" in output)

    def test_os_version_endpoint(self):
        """Test os version endpoint."""
        self.NOFLAGCMD = "-os-version"
        exit_code, output = self.container.exec_run(
            cmd=f"{self.BINARY} {self.NOFLAGCMD}", tty=True
        )

        self.assertEqual(exit_code, 0)
        self.assertTrue(b"Ubuntu" in output)

    def test_conf_endpoint(self):
        """Test conf endpoint."""
        self.NOFLAGCMD = "-show-config -conf=/etc/qat-computer/conf/conf_docker.yaml"
        exit_code, output = self.container.exec_run(
            cmd=f"{self.BINARY} {self.NOFLAGCMD}", tty=True
        )

        self.assertEqual(exit_code, 0)
        self.assertTrue(
            b'"ConfPath":"/etc/qat-computer/conf/conf_docker.yaml"' in output
        )

    def tearDown(self) -> None:
        """TearDown container object."""
        self.container.remove(force=True)
