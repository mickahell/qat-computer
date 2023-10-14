.. image:: /images/DockerGopher.png
    Docker & Go image by github.com/ashleymcnamara/gophers

QatComputer is base docker image for Qiskit.

The source code to the project is available `on GitHub <https://github.com/mickahell/qat-computer>`_.

------------

**Quickstart**

Step 0: install package

.. code-block::
   :caption: docker pull

      docker pull qatcomputer


Step 1: run experiment

.. code-block:: bash
   :caption: bash.sh

    docker run --rm \
		--name qatcomputer \
		-v ./qat_computer/resources/conf:/etc/qat-computer/conf \
		-v ./qat_computer/resources/compute:/etc/qat-computer/compute \
		qatcomputer:latest \
		compute -conf=/etc/qat-computer/conf/conf_docker.yaml


------------

**Compute**

Blablabla
Blabla

------------

**Content**

.. toctree::
  :maxdepth: 2

  Documentation Home <self>

.. toctree::
  :maxdepth: 2

  Guides <guides/index>
  API References <apidocs/index>

.. Hiding - Indices and tables
   :ref:`genindex`
   :ref:`modindex`
   :ref:`search`