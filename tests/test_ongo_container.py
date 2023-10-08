"""Tests for on go container."""
import os
from unittest import TestCase
from testcontainers.compose import DockerCompose

from .test_basic_container import call_container


class TestOnGoContainer(TestCase):
    """TestOnGoContainer."""

    def setUp(self) -> None:
        """SetUp on the go container object."""
        self.current_directory = os.path.dirname(os.path.abspath(__file__))
        self.build_img = False

    def test_full_endpoint(self):
        """Test full endpoint."""
        os.environ["CMD"] = "-conf /etc/qat-computer/conf/conf_docker.yaml"
        stdout = call_container(filepath=os.path.join(self.current_directory, "../"), waitfor="# End")

        self.assertTrue(b"ERROR" not in stdout)
        self.assertTrue(b"WARNING" not in stdout)
        self.assertTrue(b'message="4.0"' in stdout)
