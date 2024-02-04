#.. image:: /images/DockerGopher.png

#*Docker & Go image by github.com/ashleymcnamara/gophers*
*QatComputer*

QatComputer is base docker image for Qiskit.

The source code to the project is available `on GitHub <https://github.com/mickahell/qat-computer>`_.

------------

**Quickstart**

Step 0: install package

.. code-block:: bash
   :caption: docker_pull

      docker pull qatcomputer


Step 1: start container

.. code-block:: bash
   :caption: docker_run

      docker run -d --rm --privileged --cgroupns=host \
        --name qatcomputer \
        -v /sys/fs/cgroup:/sys/fs/cgroup:rw \
        -v $(PWD)/qat_computer/resources/conf:/etc/qat-computer/conf \
        -v $(PWD)/qat_computer/resources/compute:/etc/qat-computer/compute \
        qatcomputer:latest

Step 2: run experiment

.. code-block:: bash
   :caption docker_exec

      docker exec --tty qatcomputer \
	      qat-computer compute -conf=conf_docker.yaml

------------

**Content**

.. toctree::
  :maxdepth: 2

  Documentation Home <self>

.. toctree::
  :maxdepth: 2

  Guides <guides/index>

.. Hiding - Indices and tables
   :ref:`genindex`
   :ref:`modindex`
   :ref:`search`