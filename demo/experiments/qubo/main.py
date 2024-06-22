from qiskit_algorithms.utils import algorithm_globals
from qiskit_algorithms.minimum_eigensolvers import QAOA, NumPyMinimumEigensolver
from qiskit_algorithms.optimizers import COBYLA
from qiskit.primitives import Sampler
from qiskit_optimization.algorithms import (
    GroverOptimizer,
    MinimumEigenOptimizer,
)
from qiskit_optimization import QuadraticProgram
from itertools import combinations
from typing import List, Tuple
import numpy as np
import emoji
import math

backend_sim = Sampler()
algorithm_globals.random_seed = 42

## Data
dataset = [1, 2, 3, -4]
params = ["a", "b", "c", "d"]  # 1 param = 1 data point
constant = 0
gen_coef = 1

## QuadraticProgram
qubo = QuadraticProgram(name="Get minimal value.")
for param in params:
    qubo.binary_var(param)

qubo.minimize(
    constant=constant,
    linear=dataset,
    quadratic={(x, y): gen_coef for x, y in combinations(params, 2)},
)
qubo.linear_constraint(
    name="2_minimals", linear={param: 1 for param in params}, sense="==", rhs=2
)

print(qubo.prettyprint())

## Classical
exact_mes = NumPyMinimumEigensolver()
exact = MinimumEigenOptimizer(exact_mes)
exact_result = exact.solve(qubo)
print(exact_result.prettyprint())
dataset_sort = sorted(dataset)
mini = dataset.index(dataset_sort[0])
almost_mini = dataset.index(dataset_sort[1])
if exact_result.x[mini] == 1 and exact_result.x[almost_mini] == 1:
    print(
        emoji.emojize(
            "\nSame as the theorical result. \nCongratulations :party_popper: !!"
        )
    )
else:
    print(emoji.emojize("\nNot passing :cross_mark:"))

## QAOA
qaoa_mes = QAOA(sampler=backend_sim, optimizer=COBYLA(), initial_point=[0.0, 0.0])
qaoa = MinimumEigenOptimizer(qaoa_mes)
qaoa_result = qaoa.solve(qubo)
print(qaoa_result.prettyprint())
if qaoa_result.prettyprint() == exact_result.prettyprint():
    print(
        emoji.emojize(
            "\nSame as the classical result. \nCongratulations :party_popper: !!"
        )
    )
else:
    print(emoji.emojize("\nNot passing :cross_mark:"))

## Grover
nb_qubits = math.ceil(math.log2(len(dataset)))
nb_iter = math.ceil((np.pi / len(dataset)) * (math.sqrt(2 ** len(dataset))))
grover_optimizer = GroverOptimizer(
    nb_qubits, num_iterations=nb_iter, sampler=backend_sim
)
grover_result = grover_optimizer.solve(qubo)
print(grover_result.prettyprint())
if grover_result.prettyprint() == exact_result.prettyprint():
    print(
        emoji.emojize(
            "\nSame as the classical result. \nCongratulations :party_popper: !!"
        )
    )
else:
    print(emoji.emojize("\nNot passing :cross_mark:"))
if grover_result.status.value == 2:
    print(emoji.emojize(":warning: seems not usable for most of the QUBO problems"))
