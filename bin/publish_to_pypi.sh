#!/bin/bash -e

# move the dcos package
cd /dcos-cli

# move generated pypirc configuration to correct location
mv .pypirc ~/.pypirc

make clean env
source env/bin/activate
"$BASEDIR/env/bin/python" setup.py bdist_wheel upload || exit $?
echo "Wheel should now be online at: https://pypi.python.org/pypi/dcos"
deactivate

# Move down to the dcoscli package
cd cli

make clean env
source env/bin/activate
"$BASEDIR/env/bin/python" setup.py bdist_wheel upload || exit $?
echo "Wheel should now be online at: https://pypi.python.org/pypi/dcoscli"
deactivate
