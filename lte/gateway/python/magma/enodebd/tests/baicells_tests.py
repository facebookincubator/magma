"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

# pylint: disable=protected-access
from unittest import TestCase
from magma.enodebd.devices.device_utils import EnodebDeviceName
from magma.enodebd.tr069 import models
from magma.enodebd.tests.test_utils.tr069_msg_builder import \
    Tr069MessageBuilder
from magma.enodebd.tests.test_utils.enb_acs_builder import \
    EnodebAcsStateMachineBuilder


class BaicellsHandlerTests(TestCase):
    def test_initial_enb_bootup(self) -> None:
        """
        Baicells does not support configuration during initial bootup of
        eNB device. This is because it is in a REM process, and we just need
        to wait for this process to finish, ~10 minutes. Attempting to
        configure the device during this period will cause undefined
        behavior.
        As a result of this, end any provisoning sessions, which we can do
        by just sending empty HTTP responses, not even using an
        InformResponse.
        """
        acs_state_machine = \
            EnodebAcsStateMachineBuilder \
                .build_acs_state_machine(EnodebDeviceName.BAICELLS)

        # Send an Inform message
        inform_msg = Tr069MessageBuilder.get_inform('48BF74',
                                                    'BaiBS_RTS_3.1.6',
                                                    '120200002618AGP0003',
                                                    ['1 BOOT'])
        resp = acs_state_machine.handle_tr069_message(inform_msg)

        self.assertTrue(isinstance(resp, models.DummyInput),
                        'Should respond with an InformResponse')

    def test_provision_without_invasive_changes(self) -> None:
        """
        Test the scenario where:
        - eNodeB has already been powered for 10 minutes without configuration
        - Setting parameters which are 'non-invasive' on the eNodeB

        'Invasive' parameters are those which require special behavior to apply
        the changes for the eNodeB.
        """
        acs_state_machine = \
            EnodebAcsStateMachineBuilder \
                .build_acs_state_machine(EnodebDeviceName.BAICELLS)

        # Send an Inform message, wait for an InformResponse
        inform_msg = Tr069MessageBuilder.get_inform()
        resp = acs_state_machine.handle_tr069_message(inform_msg)
        self.assertTrue(isinstance(resp, models.InformResponse),
                        'Should respond with an InformResponse')

        # Send an empty http request to kick off the rest of provisioning
        req = models.DummyInput()
        resp = acs_state_machine.handle_tr069_message(req)

        # Expect a request for an optional parameter, three times
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')
        req = Tr069MessageBuilder.get_fault()
        resp = acs_state_machine.handle_tr069_message(req)
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')
        req = Tr069MessageBuilder.get_fault()
        resp = acs_state_machine.handle_tr069_message(req)
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')
        req = Tr069MessageBuilder.get_fault()
        resp = acs_state_machine.handle_tr069_message(req)

        # Expect a request for read-only params
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')
        req = Tr069MessageBuilder.get_read_only_param_values_response()

        # Send back some typical values
        # And then SM should request regular parameter values
        resp = acs_state_machine.handle_tr069_message(req)
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')

        # Send back typical values for the regular parameters
        req = Tr069MessageBuilder.\
            get_regular_param_values_response(admin_state=False,
                                              earfcndl=39150)
        resp = acs_state_machine.handle_tr069_message(req)

        # SM will be requesting object parameter values
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting object param vals')

        # Send back some typical values for object parameters
        req = Tr069MessageBuilder.get_object_param_values_response()
        resp = acs_state_machine.handle_tr069_message(req)

        # In this scenario, the ACS and thus state machine will not need
        # to delete or add objects to the eNB configuration.
        # SM should then just be attempting to set parameter values
        self.assertTrue(isinstance(resp, models.SetParameterValues),
                        'State machine should be setting param values')

        # Send back confirmation that the parameters were successfully set
        req = models.SetParameterValuesResponse()
        req.Status = 0
        resp = acs_state_machine.handle_tr069_message(req)

        # Expect a request for read-only params
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')
        req = Tr069MessageBuilder.get_read_only_param_values_response()

        # Send back some typical values
        # And then SM should continue polling the read-only params
        resp = acs_state_machine.handle_tr069_message(req)
        self.assertTrue(isinstance(resp, models.DummyInput),
                        'State machine should be ending session')

        # If a different eNB is suddenly plugged in, or the same eNB sends a
        # new Inform, enodebd should be able to handle it.
        # Send an Inform message, wait for an InformResponse
        inform_msg = Tr069MessageBuilder.get_inform()
        resp = acs_state_machine.handle_tr069_message(inform_msg)
        self.assertTrue(isinstance(resp, models.InformResponse),
                        'Should respond with an InformResponse')

        # Send an empty http request to kick off the rest of provisioning
        req = models.DummyInput()
        resp = acs_state_machine.handle_tr069_message(req)

        # Expect a request for an optional parameter, three times
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')

    def test_reboot_after_invasive_changes(self) -> None:
        """
        Test the scenario where:
        - eNodeB has already been powered for 10 minutes without configuration
        - Setting parameters which are 'invasive' on the eNodeB
        - Simulate the scenario up until reboot, and test that enodebd does
          not try to complete configuration after reboot, because it is
          waiting for REM process to finish running
        - This test does not wait the ten minutes to simulate REM process
          finishing on the Baicells eNodeB

        'Invasive' parameters are those which require special behavior to apply
        the changes for the eNodeB.

        In the case of the Baicells eNodeB, properly applying changes to
        invasive parameters requires rebooting the device.
        """
        acs_state_machine = \
            EnodebAcsStateMachineBuilder\
            .build_acs_state_machine(EnodebDeviceName.BAICELLS)

        # Send an Inform message, wait for an InformResponse
        inform_msg = Tr069MessageBuilder.get_inform('48BF74',
                                                    'BaiBS_RTS_3.1.6',
                                                    '120200002618AGP0003',
                                                    ['2 PERIODIC'])
        resp = acs_state_machine.handle_tr069_message(inform_msg)
        self.assertTrue(isinstance(resp, models.InformResponse),
                        'Should respond with an InformResponse')

        # Send an empty http request to kick off the rest of provisioning
        req = models.DummyInput()
        resp = acs_state_machine.handle_tr069_message(req)

        # Expect a request for an optional parameter, three times
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')
        req = Tr069MessageBuilder.get_fault()
        resp = acs_state_machine.handle_tr069_message(req)
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')
        req = Tr069MessageBuilder.get_fault()
        resp = acs_state_machine.handle_tr069_message(req)
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')
        req = Tr069MessageBuilder.get_fault()
        resp = acs_state_machine.handle_tr069_message(req)

        # Expect a request for read-only params
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')
        req = Tr069MessageBuilder.get_read_only_param_values_response()

        # Send back some typical values
        # And then SM should request regular parameter values
        resp = acs_state_machine.handle_tr069_message(req)
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting param values')

        # Send back typical values for the regular parameters
        req = Tr069MessageBuilder.get_regular_param_values_response()
        resp = acs_state_machine.handle_tr069_message(req)

        # SM will be requesting object parameter values
        self.assertTrue(isinstance(resp, models.GetParameterValues),
                        'State machine should be requesting object param vals')

        # Send back some typical values for object parameters
        req = Tr069MessageBuilder.get_object_param_values_response()
        resp = acs_state_machine.handle_tr069_message(req)

        # In this scenario, the ACS and thus state machine will not need
        # to delete or add objects to the eNB configuration.
        # SM should then just be attempting to set parameter values
        self.assertTrue(isinstance(resp, models.SetParameterValues),
                        'State machine should be setting param values')

        # Send back confirmation that the parameters were successfully set
        req = models.SetParameterValuesResponse()
        req.Status = 0
        resp = acs_state_machine.handle_tr069_message(req)

        # Since invasive parameters have been set, then to apply the changes
        # to the Baicells eNodeB, we need to reboot the device
        self.assertTrue(isinstance(resp, models.Reboot))
        req = Tr069MessageBuilder.get_reboot_response()
        resp = acs_state_machine.handle_tr069_message(req)

        # After the reboot has been received, enodebd should end the
        # provisioning session
        self.assertTrue(isinstance(resp, models.DummyInput),
                        'After sending command to reboot the Baicells eNodeB, '
                        'enodeb should end the TR-069 session.')

        # At this point, sometime after the eNodeB reboots, we expect it to
        # send an Inform indicating reboot. Since it should be in REM process,
        # we hold off on finishing configuration, and end TR-069 sessions.
        req = Tr069MessageBuilder.get_inform('48BF74', 'BaiBS_RTS_3.1.6',
                                             '120200002618AGP0003',
                                             ['1 BOOT', 'M Reboot'])
        resp = acs_state_machine.handle_tr069_message(req)
        self.assertTrue(isinstance(resp, models.DummyInput),
                        'After receiving a post-reboot Inform, enodebd '
                        'should end TR-069 sessions for 10 minutes to wait '
                        'for REM process to finish.')
