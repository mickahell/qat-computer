"""Tests for basic container."""
import os
from unittest import TestCase
from testcontainers.compose import DockerCompose
from testcontainers.core.waiting_utils import wait_for_logs


def call_container(
    filepath: str, waitfor: str = "", build: bool = False
) -> DockerCompose:
    """Create compose container function."""
    with DockerCompose(
        filepath=filepath, compose_file_name="docker-compose.yml", build=build
    ) as compose:
        wait_for_logs(compose, waitfor)
        stdout, _ = compose.get_logs()
    return stdout


class TestBasicContainer(TestCase):
    """TestBasicContainer."""

    def setUp(self) -> None:
        """SetUp container object."""
        self.current_directory = os.path.dirname(os.path.abspath(__file__))
        self.build_img = False

    def test_version_endpoint(self):
        """Test version endpoint."""
        os.environ["SUBCMD"] = ""
        os.environ["CMD"] = "-version"
        stdout = call_container(filepath=os.path.join(self.current_directory, "../"))

        with open(
            os.path.join(self.current_directory, "../VERSION.txt"),
            "r",
            encoding="utf-8",
        ) as vers_file:
            self.assertTrue(bytes(vers_file.read(), "utf-8") in stdout)

    def test_qiskit_version_endpoint(self):
        """Test qiskit version endpoint."""
        os.environ["SUBCMD"] = ""
        os.environ["CMD"] = "-qiskit-version"
        stdout = call_container(filepath=os.path.join(self.current_directory, "../"))

        self.assertTrue(
            b'qiskit' in stdout
        )

    def test_os_version_endpoint(self):
        """Test os version endpoint."""
        os.environ["SUBCMD"] = ""
        os.environ["CMD"] = "-os-version"
        stdout = call_container(filepath=os.path.join(self.current_directory, "../"))

        self.assertTrue(
            b'Ubuntu' in stdout
        )
