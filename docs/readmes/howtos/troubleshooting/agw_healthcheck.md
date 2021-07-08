---
id: agw_healthcheck
title: Perform Access Gateway health check
hide_title: true
---
# Perform Access Gateway health check

Performing a timely AGW health check is essential to confirm that an AGW is in good operating state, has no failures or errors and to proactively resolve any occurring issues. It is recommended to perform this operation (in addition to regular checks) before and after any changes in node (and compare) to make sure there are no undesired effects of any activity.

Affected components: AGW, enodeb, Orchestrator

## Connectivity

1.  Login to the AGW by running below the command in terminal:

    `ssh magma@<IP of AGW>`

2.  Checking Magma interfaces. Make sure eth0 and eth1 are UP.

    `ip addr`

## Enodeb Connection


1.  Check S1 and SGi interfaces can ping eNodeB(s) and internet respectively.

    `ping google.com -I eth0`

    `ping <enodeB IP> -I eth1`

2.  For managed eNB check status of eNodeB(s) attached to gateway using the cli(skip this step for unmanaged eNB):

    `sudo enodebd_cli.py get_all_status`

    An eNodeB in good state, looks similar to the below:

    ```
    magma@magma:~$ enodebd_cli.py get_all_status
    --- eNodeB Serial: 120200004917CNJ0028 ---
    IP Address..................10.0.2.243
    eNodeB connected....................ON
    eNodeB Configured...................ON
    Opstate Enabled.....................ON
    RF TX on............................ON
    RF TX desired.......................ON
    GPS Connected.......................ON
    PTP Connected......................OFF
    MME Connected.......................ON
    GPS Longitude..............-106.347936
    GPS Latitude.................35.608135
    FSM State...............Completed provisioning eNB. Awaiting new Inform.
    ```

3.  Check eNodeB at SCTP level by taking a TCP dump. There should be a heartbeat messaging between eNB and AGW IP.

    `sudo tcpdump -i any sctp`

    A sctp association in good state looks similar as below:

    ```
    magma@magma:~$ sudo tcpdump -i any sctp
    tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
    listening on any, link-type LINUX_SLL (Linux cooked), capture size 262144 bytes
    06:59:06.045369 IP 10.0.2.243.36412 > 10.0.2.242.36412: sctp (1) [HB REQ]
    06:59:06.045521 IP 10.0.2.242.36412 > 10.0.2.243.36412: sctp (1) [HB ACK]
    06:59:07.534188 IP 10.0.2.242.36412 > 10.0.2.243.36412: sctp (1) [HB REQ]
    06:59:07.544183 IP 10.0.2.243.36412 > 10.0.2.242.36412: sctp (1) [HB ACK]
    ```

## Magma Services


1.  Check all gateway services and their status by running below command. Make sure that all services are in “active (running)” state and there are no errors in any service.

    `sudo service magma@* status`

    Consider the "active (running)" duration aligns with the AGW being running, this will give you an idea of unexpected restart of services. Example, service running for `3 min and 35s`.

    ```sh
    ● magma@magmad.service - Magma magmad service
    Loaded: loaded (/etc/systemd/system/magma@magmad.service; disabled; vendor preset: enabled)
    Active: active (running) since Mon 2021-05-10 22:24:09 UTC; 3min 35s ago
    ```


2.  Check the status of OVS module with `sudo ovs-vsctl show`. Make sure the “is_connected” states are “true”.

    OVS in good state looks similar to the below:

    ```
    magma-dev:~$ sudo ovs-vsctl show
    e2bf2cb0-7bbe-48ef-a489-3341731685e1
        Manager "ptcp:6640"
        Bridge "uplink_br0"
            Port "uplink_br0"
                Interface "uplink_br0"
                    type: internal
            Port patch-agw
                Interface patch-agw
                    type: patch
                    options: {peer=patch-up}
            Port "dhcp0"
                Interface "dhcp0"
                    type: internal
        Bridge "gtp_br0"
            Controller "tcp:127.0.0.1:6633"
                is_connected: true
    ```



3.  Check connectivity with Orchestrator by running below command: `checkin_cli.py`

    An AGW connection with Orc8r in good state, looks similar to the below:

    ```sh
    magma@magma:~$ checkin_cli.py

    1. -- Testing TCP connection to controller.magma.test.io:443 --
    2. -- Testing Certificate --
    3. -- Testing SSL --
    4. -- Creating direct cloud checkin --
    5. -- Creating proxy cloud checkin --
    Success!
    ```
    Note: Verify if AGW was successfully checkin in NMS(Show as "Good" Health)

## Subscribers

1.  Check subscribers attached using the below command.

    `sudo mobility_cli.py get_subscriber_table`

2. Check if subscribers are not dropping packets due to Magma. Follow the tool https://magma.github.io/magma/docs/next/lte/debug_dp_probe


## Performance


1.  Check CPU utilization. `top`. If it is high, check which process is utilizing CPU more from output of the same command. All processes are listed there.


2.  Check memory utilization by running the same command as above. You can also verify by service using the command
    `ps -o pid,user,%mem,command ax | sort -b -k3 -r`


## Metrics

Login to NMS UI. From the left side menu options, select  “Metrics”.  Check various metrics that are available. Look  for any sudden spike or degradation that may indicate issues with the system.

  - Number of Connected eNBs (Grafana -> Dashboards -> Networks)
  - Network of Connected UE (Grafana -> Dashboards -> Networks)
  - Network of Registered UE (Grafana -> Dashboards -> Networks)
  - Attach/ Reg attempts (Grafana -> Dashboards -> Networks)
  - Attach Success Rate (Grafana -> Dashboards -> Networks)
  - S6a Authentication Success Rate (Grafana -> Dashboards -> Networks)
  - Service Request Success Rate (Grafana -> Dashboards -> Networks)
  - Session Create Success Rate (Grafana -> Dashboards -> Networks)
  - Upload/Download Throughput (Grafana -> Dashboards -> Gateway)

Note: Number of sites(enodeb) down, users affected, and outage duration are key indicators of service impact.


## Optional Features

Make sure you test any other feature that is applicable to your network

-   X2 Handover
-   S1-Flex
-   Inbound Roaming
-   External DHCP
-   UE Bridge Mode
