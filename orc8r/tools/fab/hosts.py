"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""
from fabric.api import env, lcd, local

from tools.fab import vagrant


def split_hoststring(hoststring):
    """
    Splits a host string into its user, hostname, and port components

    e.g. 'vagrant@localhost:22' -> ('vagrant', 'localhost', '22')
    """
    user = hoststring[0:hoststring.find('@')]
    ip = hoststring[hoststring.find('@') + 1:hoststring.find(':')]
    port = hoststring[hoststring.find(':') + 1:]
    return (user, ip, port)


def vagrant_setup(host, destroy_vm):
    """
    Setup the specified vagrant box

    host: the Vagrant box to setup, e.g. "magma"
    """
    if destroy_vm:
        vagrant.teardown_vagrant(host)
    vagrant.setup_env_vagrant(host)
    return env.hosts[0]


def ansible_setup(hoststr, ansible_group, playbook,
                  preburn='false', full_provision='true'):
    """
    Setup the specified ansible machine

    hoststr: the host string of the target host
             e.g. vagrant@192.168.60.10:22

    ansible_group: The group the deploy targets
             e.g. "dev"

    preburn: 'true' to run preburn tasks, 'false' to skip them.
             Defaults to 'false'

    full_provision: 'true' to run post-preburn tasks, 'false' to skip them.
                    Defaults to 'true'
    """
    env.hosts = [hoststr]
    # Provision the gateway host
    (user, ip, port) = split_hoststring(hoststr)

    local("echo '[%s]\nhost ansible_host=%s ansible_user=%s"
          " ansible_port=%s' > /tmp/hosts" % (ansible_group, ip, user, port))
    local("ansible-playbook -i /tmp/hosts deploy/%s "
          "--extra-vars '{\"preburn\": %s, \"full_provision\": %s}'" %
          (playbook, preburn, full_provision))
