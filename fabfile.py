from fabric.api import *
from fabric.colors import *
from fabric.contrib.console import confirm

import logging
logging.basicConfig()

env.user = 'cydev'
env.base_dir = '/cydev'
env.use_ssh_config = True
env.hosts = ['cydev.ru']


def deploy(branch='master', restart='yes'):
    with cd(env.base_dir):
        run('git fetch --all')
        run('git reset --hard origin/master')
