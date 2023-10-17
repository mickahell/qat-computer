"""Tests for compute subcmd container."""
import os
import docker
from unittest import TestCase

from .test_basic_container import run_container, build_image


class TestComputeContainer(TestCase):
    """TestComputeContainer."""

    def setUp(self) -> None:
        """SetUp compute container object."""
        self.current_directory = os.path.dirname(os.path.abspath(__file__))
        self.build_img = False
        self.BINARY = "qat-computer"
        self.SUBCMD = "compute"
        self.FLAGCMD = ""
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

    def test_full_endpoint(self):
        """Test full compute endpoint."""
        self.FLAGCMD = "-conf /etc/qat-computer/conf/conf_docker.yaml"
        exit_code, output = self.container.exec_run(
            cmd=f"{self.BINARY} {self.SUBCMD} {self.FLAGCMD}", tty=True
        )

        self.assertEqual(exit_code, 0)
        self.assertTrue(b"ERROR" not in output)
        self.assertTrue(b"WARNING" not in output)
        self.assertTrue(b'message="4.0"' in output)

    def tearDown(self) -> None:
        """TearDown compute container object."""
        self.container.remove(force=True)
