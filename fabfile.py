from fabric.api import *
from fabric.colors import *
from fabric.contrib.console import confirm

import logging
logging.basicConfig()

env.user = 'cydev'
env.base_dir = '/home/cydev'
env.use_ssh_config = True
env.hosts = ['cydev.ru']


def deploy(branch='master', restart='yes'):
    with cd(env.base_dir):
        run('git fetch')
        run('git checkout %s' % branch)
        run('git pull --rebase origin %s' % branch)

def restart():
    env.user = 'root'
    run('systemctl restart cydev')


def status():
    env.user = 'root'
    run('systemctl status cydev')

