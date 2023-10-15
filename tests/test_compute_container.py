"""Tests for on go container."""
import os
from unittest import TestCase

from .test_basic_container import call_container


class TestComputeContainer(TestCase):
    """TestComputeContainer."""

    def setUp(self) -> None:
        """SetUp compute container object."""
        self.current_directory = os.path.dirname(os.path.abspath(__file__))
        self.build_img = False
    
    def test_conf_endpoint(self):
        """Test conf endpoint."""
        os.environ["SUBCMD"] = "compute"
        os.environ["CMD"] = "-conf /etc/qat-computer/conf/conf_docker.yaml -show-config"
        stdout = call_container(filepath=os.path.join(self.current_directory, "../"))

        self.assertTrue(
            b'"ConfPath":"/etc/qat-computer/conf/conf_docker.yaml"' in stdout
        )

    def test_full_endpoint(self):
        """Test full compute endpoint."""
        os.environ["SUBCMD"] = "compute"
        os.environ["CMD"] = "-conf /etc/qat-computer/conf/conf_docker.yaml"
        stdout = call_container(
            filepath=os.path.join(self.current_directory, "../"), waitfor="# End"
        )

        self.assertTrue(b"ERROR" not in stdout)
        self.assertTrue(b"WARNING" not in stdout)
        self.assertTrue(b'message="4.0"' in stdout)
