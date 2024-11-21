"""Tests for compute subcmd container."""

import os
from unittest import TestCase
import docker

from .test_basic_container import run_container


class TestComputeContainer(TestCase):
    """TestComputeContainer."""

    def setUp(self) -> None:
        """SetUp compute container object."""
        self.current_directory = os.path.dirname(os.path.abspath(__file__))
        self.build_img = False
        self.binary = "qat-computer"
        self.subcmd = "compute"
        self.flagcmd = ""
        self.client = docker.from_env()
        self.container = run_container(
            docker_env=self.client, current_directory=self.current_directory
        )

    def test_full_endpoint(self):
        """Test full compute endpoint."""
        self.flagcmd = "-conf conf_docker.yaml"
        exit_code, output = self.container.exec_run(
            cmd=f"{self.binary} {self.subcmd} {self.flagcmd}", tty=True
        )

        self.assertEqual(exit_code, 0)
        self.assertTrue(b"ERROR" not in output)
        self.assertTrue(b"WARNING" not in output)
        self.assertTrue(b'message="4.0"' in output)

    def tearDown(self) -> None:
        """TearDown compute container object."""
        self.container.remove(force=True)
