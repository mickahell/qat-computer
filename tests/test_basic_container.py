"""Tests for basic container."""
import os
from unittest import TestCase
from testcontainers.compose import DockerCompose


class TestContainer(TestCase):
    """TestContainer."""

    def setUp(self) -> None:
        """SetUp container object."""
        self.current_directory = os.path.dirname(os.path.abspath(__file__))
        self.build_img = False

    def test_version_endpoint(self):
        """Test version endpoint."""
        os.environ["CMD"] = "-version"
        with DockerCompose(
            filepath=os.path.join(self.current_directory, "../"),
            compose_file_name="docker-compose.yml",
            build=self.build_img,
        ) as compose:
            stdout, _ = compose.get_logs()

        with open(
            os.path.join(self.current_directory, "../VERSION.txt"),
            "r",
            encoding="utf-8",
        ) as vers_file:
            self.assertTrue(bytes(vers_file.read(), "utf-8") in stdout)

    def test_conf_endpoint(self):
        """Test conf endpoint."""
        os.environ["CMD"] = "-conf /etc/qat-computer/conf/conf_docker.yaml -show-config"
        with DockerCompose(
            filepath=os.path.join(self.current_directory, "../"),
            compose_file_name="docker-compose.yml",
            build=self.build_img,
        ) as compose:
            stdout, _ = compose.get_logs()

        self.assertTrue(
            b'"ConfPath":"/etc/qat-computer/conf/conf_docker.yaml"' in stdout
        )
