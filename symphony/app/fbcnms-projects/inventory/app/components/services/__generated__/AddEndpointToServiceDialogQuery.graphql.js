/**
 * @generated
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 **/

 /**
 * @flow
 * @relayHash c559ee0e1e4d58ca9d85b1c766c0e03c
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
type AvailablePortsTable_ports$ref = any;
export type FilterOperator = "CONTAINS" | "DATE_GREATER_OR_EQUAL_THAN" | "DATE_GREATER_THAN" | "DATE_LESS_OR_EQUAL_THAN" | "DATE_LESS_THAN" | "IS" | "IS_NOT_ONE_OF" | "IS_ONE_OF" | "%future added value";
export type PortFilterType = "LOCATION_INST" | "LOCATION_INST_EXTERNAL_ID" | "PORT_DEF" | "PORT_INST_EQUIPMENT" | "PORT_INST_HAS_LINK" | "PROPERTY" | "SERVICE_INST" | "%future added value";
export type PropertyKind = "bool" | "date" | "datetime_local" | "email" | "enum" | "float" | "gps_location" | "int" | "node" | "range" | "string" | "%future added value";
export type PortFilterInput = {|
  filterType: PortFilterType,
  operator: FilterOperator,
  boolValue?: ?boolean,
  stringValue?: ?string,
  propertyValue?: ?PropertyTypeInput,
  idSet?: ?$ReadOnlyArray<string>,
  stringSet?: ?$ReadOnlyArray<string>,
  maxDepth?: ?number,
|};
export type PropertyTypeInput = {|
  id?: ?string,
  externalId?: ?string,
  name: string,
  type: PropertyKind,
  nodeType?: ?string,
  index?: ?number,
  category?: ?string,
  stringValue?: ?string,
  intValue?: ?number,
  booleanValue?: ?boolean,
  floatValue?: ?number,
  latitudeValue?: ?number,
  longitudeValue?: ?number,
  rangeFromValue?: ?number,
  rangeToValue?: ?number,
  isEditable?: ?boolean,
  isInstanceProperty?: ?boolean,
  isMandatory?: ?boolean,
  isDeleted?: ?boolean,
|};
export type AddEndpointToServiceDialogQueryVariables = {|
  filters: $ReadOnlyArray<PortFilterInput>
|};
export type AddEndpointToServiceDialogQueryResponse = {|
  +equipmentPorts: {|
    +edges: $ReadOnlyArray<{|
      +node: ?{|
        +id: string,
        +definition: {|
          +id: string,
          +name: string,
        |},
        +$fragmentRefs: AvailablePortsTable_ports$ref,
      |}
    |}>
  |}
|};
export type AddEndpointToServiceDialogQuery = {|
  variables: AddEndpointToServiceDialogQueryVariables,
  response: AddEndpointToServiceDialogQueryResponse,
|};
*/


/*
query AddEndpointToServiceDialogQuery(
  $filters: [PortFilterInput!]!
) {
  equipmentPorts(filters: $filters, first: 50) {
    edges {
      node {
        id
        definition {
          id
          name
        }
        ...AvailablePortsTable_ports
      }
    }
  }
}

fragment AvailablePortsTable_ports on EquipmentPort {
  id
  parentEquipment {
    id
    name
    ...EquipmentBreadcrumbs_equipment
  }
  definition {
    id
    name
    portType {
      name
      id
    }
    visibleLabel
  }
}

fragment EquipmentBreadcrumbs_equipment on Equipment {
  id
  name
  equipmentType {
    id
    name
  }
  locationHierarchy {
    id
    name
    locationType {
      name
      id
    }
  }
  positionHierarchy {
    id
    definition {
      id
      name
      visibleLabel
    }
    parentEquipment {
      id
      name
      equipmentType {
        id
        name
      }
    }
  }
}
*/

const node/*: ConcreteRequest*/ = (function(){
var v0 = [
  {
    "kind": "LocalArgument",
    "name": "filters",
    "type": "[PortFilterInput!]!",
    "defaultValue": null
  }
],
v1 = [
  {
    "kind": "Variable",
    "name": "filters",
    "variableName": "filters"
  },
  {
    "kind": "Literal",
    "name": "first",
    "value": 50
  }
],
v2 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "id",
  "args": null,
  "storageKey": null
},
v3 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "name",
  "args": null,
  "storageKey": null
},
v4 = [
  (v2/*: any*/),
  (v3/*: any*/)
],
v5 = [
  (v3/*: any*/),
  (v2/*: any*/)
],
v6 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "visibleLabel",
  "args": null,
  "storageKey": null
},
v7 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "equipmentType",
  "storageKey": null,
  "args": null,
  "concreteType": "EquipmentType",
  "plural": false,
  "selections": (v4/*: any*/)
};
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "AddEndpointToServiceDialogQuery",
    "type": "Query",
    "metadata": null,
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "equipmentPorts",
        "storageKey": null,
        "args": (v1/*: any*/),
        "concreteType": "EquipmentPortConnection",
        "plural": false,
        "selections": [
          {
            "kind": "LinkedField",
            "alias": null,
            "name": "edges",
            "storageKey": null,
            "args": null,
            "concreteType": "EquipmentPortEdge",
            "plural": true,
            "selections": [
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "node",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentPort",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "definition",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPortDefinition",
                    "plural": false,
                    "selections": (v4/*: any*/)
                  },
                  {
                    "kind": "FragmentSpread",
                    "name": "AvailablePortsTable_ports",
                    "args": null
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  },
  "operation": {
    "kind": "Operation",
    "name": "AddEndpointToServiceDialogQuery",
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "equipmentPorts",
        "storageKey": null,
        "args": (v1/*: any*/),
        "concreteType": "EquipmentPortConnection",
        "plural": false,
        "selections": [
          {
            "kind": "LinkedField",
            "alias": null,
            "name": "edges",
            "storageKey": null,
            "args": null,
            "concreteType": "EquipmentPortEdge",
            "plural": true,
            "selections": [
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "node",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentPort",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "definition",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPortDefinition",
                    "plural": false,
                    "selections": [
                      (v2/*: any*/),
                      (v3/*: any*/),
                      {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "portType",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "EquipmentPortType",
                        "plural": false,
                        "selections": (v5/*: any*/)
                      },
                      (v6/*: any*/)
                    ]
                  },
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "parentEquipment",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Equipment",
                    "plural": false,
                    "selections": [
                      (v2/*: any*/),
                      (v3/*: any*/),
                      (v7/*: any*/),
                      {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "locationHierarchy",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "Location",
                        "plural": true,
                        "selections": [
                          (v2/*: any*/),
                          (v3/*: any*/),
                          {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "locationType",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "LocationType",
                            "plural": false,
                            "selections": (v5/*: any*/)
                          }
                        ]
                      },
                      {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "positionHierarchy",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "EquipmentPosition",
                        "plural": true,
                        "selections": [
                          (v2/*: any*/),
                          {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "definition",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "EquipmentPositionDefinition",
                            "plural": false,
                            "selections": [
                              (v2/*: any*/),
                              (v3/*: any*/),
                              (v6/*: any*/)
                            ]
                          },
                          {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "parentEquipment",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "Equipment",
                            "plural": false,
                            "selections": [
                              (v2/*: any*/),
                              (v3/*: any*/),
                              (v7/*: any*/)
                            ]
                          }
                        ]
                      }
                    ]
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  },
  "params": {
    "operationKind": "query",
    "name": "AddEndpointToServiceDialogQuery",
    "id": null,
    "text": "query AddEndpointToServiceDialogQuery(\n  $filters: [PortFilterInput!]!\n) {\n  equipmentPorts(filters: $filters, first: 50) {\n    edges {\n      node {\n        id\n        definition {\n          id\n          name\n        }\n        ...AvailablePortsTable_ports\n      }\n    }\n  }\n}\n\nfragment AvailablePortsTable_ports on EquipmentPort {\n  id\n  parentEquipment {\n    id\n    name\n    ...EquipmentBreadcrumbs_equipment\n  }\n  definition {\n    id\n    name\n    portType {\n      name\n      id\n    }\n    visibleLabel\n  }\n}\n\nfragment EquipmentBreadcrumbs_equipment on Equipment {\n  id\n  name\n  equipmentType {\n    id\n    name\n  }\n  locationHierarchy {\n    id\n    name\n    locationType {\n      name\n      id\n    }\n  }\n  positionHierarchy {\n    id\n    definition {\n      id\n      name\n      visibleLabel\n    }\n    parentEquipment {\n      id\n      name\n      equipmentType {\n        id\n        name\n      }\n    }\n  }\n}\n",
    "metadata": {}
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = '2f7b2062fec6a1371d76385835b84b32';
module.exports = node;
