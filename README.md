# cf

cf provides a change point detection algorithm called changefinder[1] in golang.
The code in this project is a rewrite version of changefinder package published at [changefinder - PiPY](https://pypi.org/project/changefinder/).

If you want to know the usage of this package, see the `cf_test.go`. You can also see the graphs which show calculated scores with the changefinder package in python and this package at `testdata/plot_graphs.ipynb`.

[1] Kenji Yamanishi and Jun-ichi Takeuchi. 2002. A unifying framework for detecting outliers and change points from non-stationary time series data. In Proceedings of the eighth ACM SIGKDD international conference on Knowledge discovery and data mining (KDD '02). ACM, New York, NY, USA, 676-681. DOI=http://dx.doi.org/10.1145/775047.775148

## Milestones
- implements ChangeFinderARIMA

## Licence
This software is provided under the MIT License, see LICENSE.
