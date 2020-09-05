/*
   Copyright 2020 The Magma Authors.
   This source code is licensed under the BSD-style license found in the
   LICENSE file in the root directory of this source tree.
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
 */
#include <bitset>
#include <sstream>
#include <cstdint>
#include <cstring>
#include "5GSMobileIdentity.h"
#include "CommonDefs.h"
using namespace std;
namespace magma5g
{

  M5GSMobileIdentityMsg::M5GSMobileIdentityMsg()
  {
  };
  M5GSMobileIdentityMsg::~M5GSMobileIdentityMsg()
  {
  };
  GutiM5GSMobileIdentity::GutiM5GSMobileIdentity()
  {
  };
  GutiM5GSMobileIdentity::~GutiM5GSMobileIdentity()
  {
  };
  ImeiM5GSMobileIdentity::ImeiM5GSMobileIdentity()
  {
  };
  ImeiM5GSMobileIdentity::~ImeiM5GSMobileIdentity()
  {
  };
  ImsiM5GSMobileIdentity::ImsiM5GSMobileIdentity()
  {
  };
  ImsiM5GSMobileIdentity::~ImsiM5GSMobileIdentity()
  {
  };
  SuciM5GSMobileIdentity::SuciM5GSMobileIdentity()
  {
  };
  SuciM5GSMobileIdentity::~SuciM5GSMobileIdentity()
  {
  };
  TmsiM5GSMobileIdentity::TmsiM5GSMobileIdentity()
  {
  };
  TmsiM5GSMobileIdentity::~TmsiM5GSMobileIdentity()
  {
  };
  M5GSMobileIdentityIe::M5GSMobileIdentityIe()
  {
  };
  M5GSMobileIdentityIe::~M5GSMobileIdentityIe()
  {
  };

  // Decode GutiMobileIdentity IE Message
  int M5GSMobileIdentityMsg::DecodeGutiMobileIdentityMsg(
      GutiM5GSMobileIdentity* guti, uint8_t* buffer, uint8_t ielen) {

    int decoded = 0;

    MLOG(MDEBUG) << "         DecodeGutiMobileIdentityMsg : ";
    guti->spare = (*(buffer + decoded) >> 4) & 0xf;

    // For the GUTI, bits 5 to 8 of octet 3 are coded as "1111"
    if (guti->spare != 0xf) {
      MLOG(MERROR) << "Error: " << std::dec << TLV_VALUE_DOESNT_MATCH;
      return(TLV_VALUE_DOESNT_MATCH);
    }

    guti->oddeven        = (*(buffer + decoded) >> 3) & 0x1;
    guti->typeofidentity = *(buffer + decoded) & 0x7;

    if (guti->typeofidentity != M5GSMobileIdentityMsg_GUTI) {
      MLOG(MERROR) << "Error: " << std::dec << TLV_VALUE_DOESNT_MATCH;
      return(TLV_VALUE_DOESNT_MATCH);
    }

    decoded++;
    guti->mcc_digit2 = (*(buffer + decoded) >> 4) & 0xf;
    guti->mcc_digit1 = *(buffer + decoded) & 0xf;
    decoded++;
    guti->mnc_digit3 = (*(buffer + decoded) >> 4) & 0xf;
    guti->mcc_digit3 = *(buffer + decoded) & 0xf;
    decoded++;
    guti->mnc_digit2 = (*(buffer + decoded) >> 4) & 0xf;
    guti->mnc_digit1 = *(buffer + decoded) & 0xf;
    decoded++;
    guti->amfregionid = *(buffer + decoded);
    decoded++;
    guti->amfsetid = *(buffer + decoded);
    decoded++;
    guti->amfsetid1 = (*(buffer + decoded) >> 6) & 0x3;
    guti->amfpointer = *(buffer + decoded) & 0x3f;
    decoded++;

    guti->tmsi1 = *(buffer + decoded);
    decoded++;
    guti->tmsi2 = *(buffer + decoded);
    decoded++;
    guti->tmsi3 = *(buffer + decoded);
    decoded++;
    guti->tmsi4 = *(buffer + decoded);
    decoded++;

    MLOG(MDEBUG) << "           oddeven = " << hex << int(guti->oddeven)<<"\n";
    MLOG(MDEBUG) << "           mcc_digit2 = " << hex << int(guti->mcc_digit2)<<"\n";
    MLOG(MDEBUG) << "           mcc_digit1 = " << hex << int(guti->mcc_digit1)<<"\n";
    MLOG(MDEBUG) << "           mnc_digit3 = " << hex << int(guti->mnc_digit3)<<"\n";
    MLOG(MDEBUG) << "           mcc_digit3 = " << hex << int(guti->mcc_digit3)<<"\n";
    MLOG(MDEBUG) << "           mnc_digit2 = " << hex << int(guti->mnc_digit2)<<"\n";
    MLOG(MDEBUG) << "           mnc_digit1 = " << hex << int(guti->mnc_digit1)<<"\n";

    MLOG(MDEBUG) << "           amfregionid = " << hex << int(guti->amfregionid)<<"\n";
    MLOG(MDEBUG) << "           amfsetid = " << hex << int(guti->amfsetid)<<"\n";
    MLOG(MDEBUG) << "           amfsetid1 = " << hex << int(guti->amfsetid1)<<"\n";
    MLOG(MDEBUG) << "           amfpointer = " << hex << int(guti->amfpointer)<<"\n";
    MLOG(MDEBUG) << "           tmsi1 = " << hex << int(guti->tmsi1)<<"\n";
    MLOG(MDEBUG) << "           tmsi2 = " << hex << int(guti->tmsi2)<<"\n";
    MLOG(MDEBUG) << "           tmsi3 = " << hex << int(guti->tmsi3)<<"\n";
    MLOG(MDEBUG) << "           tmsi4 = " << hex << int(guti->tmsi4)<<"\n";


    return(decoded);
  }

  // Decode ImeiMobileIdentity IE
  int M5GSMobileIdentityMsg::DecodeImeiMobileIdentityMsg(
      ImeiM5GSMobileIdentity* imei, uint8_t* buffer, uint8_t ielen) {
    int decoded = 0;

    MLOG(MDEBUG) << "         DecodeImeiMobileIdentityMsg : "<<"\n";
    imei->identity_digit1 = (*(buffer + decoded) >> 4) & 0xf;

    if (imei->identity_digit1 != 0xf) {
      MLOG(MERROR) << "Error: " << std::hex << TLV_VALUE_DOESNT_MATCH;
      return(TLV_VALUE_DOESNT_MATCH);
    }

    imei->oddeven        = (*(buffer + decoded) >> 3) & 0x1;
    imei->typeofidentity = *(buffer + decoded) & 0x7;

    if (imei->typeofidentity != M5GSMobileIdentityMsg_IMEI) {
      MLOG(MERROR) << "Error: " << std::dec << TLV_VALUE_DOESNT_MATCH;
      return(TLV_VALUE_DOESNT_MATCH);
    }

    decoded++;
    imei->identity_digit3 = (*(buffer + decoded) >> 4) & 0xf;
    imei->identity_digit2 = *(buffer + decoded) & 0xf;

    decoded++;

    MLOG(MDEBUG) << "           oddeven = " << hex << int(imei->oddeven)<<"\n";
    MLOG(MDEBUG) << "           digit1 = " << hex << int(imei->identity_digit1)<<"\n";
    MLOG(MDEBUG) << "           digit2 = " << hex << int(imei->identity_digit2)<<"\n";
    MLOG(MDEBUG) << "           digit3 = " << hex << int(imei->identity_digit3)<<"\n";

    return(decoded);
  };

  // Decode ImsiMobileIdentity IE
  int M5GSMobileIdentityMsg::DecodeImsiMobileIdentityMsg(
      ImsiM5GSMobileIdentity* imsi, uint8_t* buffer, uint8_t ielen) {
    int decoded = 0;
    int schemeOutLen = 0;
    uint16_t i = 0;

    MLOG(MDEBUG) << "      DecodeImsiMobileIdentityMsg:"<<"\n";
    memset(&imsi->scheme_output, 0, SCHEME_OUTPUT_MAX);

    imsi->spare2 = (*(buffer + decoded) >> 7) & 0x1;

    imsi->supiformat        = (*(buffer + decoded) >> 4) & 0x7;
    imsi->spare1 = (*(buffer + decoded) >> 3) & 0x1;
    imsi->typeofidentity = *(buffer + decoded) & 0x7;

    if (imsi->typeofidentity != M5GSMobileIdentityMsg_IMSI) {
      MLOG(MERROR) << "Error: " << std::hex << TLV_VALUE_DOESNT_MATCH;
      return(TLV_VALUE_DOESNT_MATCH);
    }

    decoded++;
    imsi->mcc_digit2 = (*(buffer + decoded) >> 4) & 0xf;
    imsi->mcc_digit1 = *(buffer + decoded) & 0xf;
    decoded++;
    imsi->mnc_digit3 = (*(buffer + decoded) >> 4) & 0xf;
    imsi->mcc_digit3 = *(buffer + decoded) & 0xf;
    decoded++;
    imsi->mnc_digit2 = (*(buffer + decoded) >> 4) & 0xf;
    imsi->mnc_digit1 = *(buffer + decoded) & 0xf;
    decoded++;

    imsi->routingindicatordigit2 = (*(buffer + decoded) >> 4) & 0xf;
    imsi->routingindicatordigit1 = *(buffer + decoded) & 0xf;
    decoded++;
    imsi->routingindicatordigit4 = (*(buffer + decoded) >> 4) & 0xf;
    imsi->routingindicatordigit3 = *(buffer + decoded) & 0xf;
    decoded++;

    imsi->spare6 = (*(buffer + decoded) >> 7) & 0x1;
    imsi->spare5 = (*(buffer + decoded) >> 6) & 0x1;
    imsi->spare4 = (*(buffer + decoded) >> 5) & 0x1;
    imsi->spare3 = (*(buffer + decoded) >> 4) & 0x1;
    imsi->protect_schm_id = *(buffer + decoded) & 0xf;
    decoded++;

    imsi->home_nw_id = *(buffer + decoded);
    decoded++;

    schemeOutLen = ielen - decoded;
    if (memcpy(&imsi->scheme_output, (buffer + decoded), schemeOutLen))
      decoded = ielen;

    MLOG(MDEBUG) << "           spare2 = 0x" << hex << int(imsi->spare2)<<"\n";
    MLOG(MDEBUG) << "           supiformat = 0x" << hex << int(imsi->supiformat)<<"\n";
    MLOG(MDEBUG) << "           spare1 = 0x" << hex << int(imsi->spare1)<<"\n";
    MLOG(MDEBUG) << "           typeofidentity = 0x" << hex << int(imsi->typeofidentity)<<"\n";
    MLOG(MDEBUG) << "           mcc_digit2 = 0x" << hex << int(imsi->mcc_digit2)<<"\n";
    MLOG(MDEBUG) << "           mcc_digit1 = 0x" << hex << int(imsi->mcc_digit1)<<"\n";
    MLOG(MDEBUG) << "           mnc_digit3 = 0x" << hex << int(imsi->mnc_digit3)<<"\n";
    MLOG(MDEBUG) << "           mcc_digit3 = 0x" << hex << int(imsi->mcc_digit3)<<"\n";
    MLOG(MDEBUG) << "           mnc_digit2 = 0x" << hex << int(imsi->mnc_digit2)<<"\n";
    MLOG(MDEBUG) << "           mnc_digit1 = 0x" << hex << int(imsi->mnc_digit1)<<"\n";
    MLOG(MDEBUG) << "           routingindicatordigit2 = 0x" << hex << int(imsi->routingindicatordigit2)<<"\n";
    MLOG(MDEBUG) << "           routingindicatordigit1 = 0x" << hex << int(imsi->routingindicatordigit1)<<"\n";
    MLOG(MDEBUG) << "           routingindicatordigit4 = 0x" << hex << int(imsi->routingindicatordigit4)<<"\n";
    MLOG(MDEBUG) << "           routingindicatordigit3 = 0x" << hex << int(imsi->routingindicatordigit3)<<"\n";
    MLOG(MDEBUG) << "           spare6 = 0x" << hex << int(imsi->spare6)<<"\n";
    MLOG(MDEBUG) << "           spare5 = 0x" << hex << int(imsi->spare5)<<"\n";
    MLOG(MDEBUG) << "           spare4 = 0x" << hex << int(imsi->spare4)<<"\n";
    MLOG(MDEBUG) << "           spare3 = 0x" << hex << int(imsi->spare3)<<"\n";
    MLOG(MDEBUG) << "           protect_schm_id = 0x" << hex << int(imsi->protect_schm_id)<<"\n";
    MLOG(MDEBUG) << "           home_nw_id = 0x" << hex << int(imsi->home_nw_id)<<"\n";
    MLOG(MDEBUG) << "           scheme_output = ";

    for(i; i < schemeOutLen; i++) {
       MLOG(MDEBUG) << "  0x" << hex << (int)(imsi->scheme_output[i]);
    }
    MLOG(MDEBUG) << endl;

    return(decoded);
  };

  // TBD
  // Decode SuciMobileIdentity IE 
  int M5GSMobileIdentityMsg::DecodeSuciMobileIdentityMsg(
      SuciM5GSMobileIdentity* suci, uint8_t* buffer, uint8_t ielen) {
    int decoded = 0;

    MLOG(MDEBUG) << "         DecodeSuciMobileIdentityMsg:"<<"\n";
    suci->spare2 = (*(buffer + decoded) >> 7) & 0x1;
    suci->supiformat        = (*(buffer + decoded) >> 4) & 0x7;
    suci->spare1 = (*(buffer + decoded) >> 3) & 0x1;

    suci->typeofidentity = *(buffer + decoded) & 0x7;

    if (suci->typeofidentity != M5GSMobileIdentityMsg_SUCI) {
      MLOG(MDEBUG) << "TLV_VALUE_DOESNT_MATCH error";
      return(TLV_VALUE_DOESNT_MATCH);
    }

    decoded++;

    //TBD
    suci->sucinai = *(buffer + decoded);
    decoded++;


    decoded++;

    return(decoded);
  };

  // Decode TmsiMobileIdentity IE
  int M5GSMobileIdentityMsg::DecodeTmsiMobileIdentityMsg (
      TmsiM5GSMobileIdentity* tmsi, uint8_t* buffer, uint8_t ielen) {
    int decoded = 0;
    
    MLOG(MDEBUG) << "         DecodeTmsiMobileIdentityMsg:"<<"\n";
    tmsi->spare = (*(buffer + decoded) >> 4) & 0xf;

    if (tmsi->spare != 0xf) {
      MLOG(MDEBUG) << "Error: " << int(TLV_VALUE_DOESNT_MATCH);
      return(TLV_VALUE_DOESNT_MATCH);
    }

    tmsi->oddeven        = (*(buffer + decoded) >> 3) & 0x1;
    tmsi->typeofidentity = *(buffer + decoded) & 0x7;

    if (tmsi->typeofidentity != M5GSMobileIdentityMsg_TMSI) {
      MLOG(MDEBUG) << "Error: " << int(TLV_VALUE_DOESNT_MATCH);
      return(TLV_VALUE_DOESNT_MATCH);
    }

    decoded++;

    tmsi->amfsetid = *(buffer + decoded);
    decoded++;
    tmsi->amfsetid1 = (*(buffer + decoded) >> 6) & 0x2;
    tmsi->amfpointer = *(buffer + decoded) & 0x3f;
    decoded++;

    tmsi->m5gtmsi1 = *(buffer + decoded);
    decoded++;
    tmsi->m5gtmsi2 = *(buffer + decoded);
    decoded++;
    tmsi->m5gtmsi3 = *(buffer + decoded);
    decoded++;
    tmsi->m5gtmsi4 = *(buffer + decoded);
    decoded++;

    MLOG(MDEBUG) << "           spare2 = " << hex << int(tmsi->spare)<<"\n";
    MLOG(MDEBUG) << "           oddeven = " << hex << int(tmsi->oddeven)<<"\n";
    MLOG(MDEBUG) << "           typeofidentity = " << hex << int(tmsi->typeofidentity)<<"\n";
    MLOG(MDEBUG) << "           amfsetid = " << hex << int(tmsi->amfsetid)<<"\n";
    MLOG(MDEBUG) << "           amfsetid1 = " << hex << int(tmsi->amfsetid1)<<"\n";
    MLOG(MDEBUG) << "           amfpointer = " << hex << int(tmsi->amfpointer)<<"\n";
    MLOG(MDEBUG) << "           m5gtmsi1 = " << hex << int(tmsi->m5gtmsi1)<<"\n";
    MLOG(MDEBUG) << "           m5gtmsi2 = " << hex << int(tmsi->m5gtmsi2)<<"\n";
    MLOG(MDEBUG) << "           m5gtmsi3 = " << hex << int(tmsi->m5gtmsi3)<<"\n";
    MLOG(MDEBUG) << "           m5gtmsi4 = " << hex << int(tmsi->m5gtmsi4)<<"\n";

    return(decoded);
  };

  // Decode M5GSMobileIdentity IE
  int M5GSMobileIdentityMsg::DecodeM5GSMobileIdentityMsg(M5GSMobileIdentityMsg *mg5smobileidentity, uint8_t iei, uint8_t *buffer, uint32_t len) 
  {
    int decoded_rc = TLV_VALUE_DOESNT_MATCH;
    int decoded    = 0;
    uint16_t ielen  = 0;

    MLOG(MDEBUG) << "    DecodeM5GSMobileIdentityMsg : "<<"\n";
    if (iei > 0) {
      CHECK_IEI_DECODER(iei, (unsigned char)*buffer);
      decoded++;
    }
    IES_DECODE_U16(buffer, decoded, ielen);
    CHECK_LENGTH_DECODER(len - decoded, ielen);
    unsigned char typeofidentity = *(buffer + decoded) & 0x7;

    MLOG(MDEBUG) << "      typeofid = 0x" << hex << bitset<4>(int(typeofidentity)) << " ielen = 0x"<< hex << bitset<16>(int(ielen))<<"\n";

    if (typeofidentity == M5GSMobileIdentityMsg_SUCI) {
      decoded_rc = DecodeSuciMobileIdentityMsg(
          &mg5smobileidentity->mobileidentity.suci, buffer, ielen);
      MLOG(MDEBUG) << "Type suci";
    } else if (typeofidentity == M5GSMobileIdentityMsg_GUTI) {
      decoded_rc = DecodeGutiMobileIdentityMsg(
          &mg5smobileidentity->mobileidentity.guti, buffer + decoded, ielen);
      MLOG(MDEBUG) << "Type guti";
    } else if (typeofidentity == M5GSMobileIdentityMsg_IMEI) {
      decoded_rc = DecodeImeiMobileIdentityMsg(
          &mg5smobileidentity->mobileidentity.imei, buffer + decoded, ielen);
      MLOG(MDEBUG) << "Type imei";
    } else if (typeofidentity == M5GSMobileIdentityMsg_TMSI) {
      decoded_rc = DecodeTmsiMobileIdentityMsg(
          &mg5smobileidentity->mobileidentity.tmsi, buffer + decoded, ielen);
      MLOG(MDEBUG) << "Type tmsi";
    } else if (typeofidentity == M5GSMobileIdentityMsg_IMSI) {
      decoded_rc = DecodeImsiMobileIdentityMsg(
          &mg5smobileidentity->mobileidentity.imsi, buffer + decoded, ielen);
      MLOG(MDEBUG) << "Type imsi";
    }

    if (decoded_rc < 0) {
      MLOG(MERROR) << "Decode Error";
      return decoded_rc;
    }
    return (decoded + decoded_rc);

  };

  // Encode GutiMobileIdentity IE
  int M5GSMobileIdentityMsg::EncodeGutiMobileIdentityMsg(
      GutiM5GSMobileIdentity* guti, uint8_t* buffer) {
    uint32_t encoded = 0;

    MLOG(MDEBUG) << "EncodeGutiMobileIdentityMsg:";
    *(buffer + encoded) =
      0xf0 | ((guti->oddeven & 0x01)<< 3) | (guti->typeofidentity & 0x7);

    MLOG(MDEBUG) << "oddeven typeofidentity = " << hex << int(*(buffer + encoded));
    encoded++;
    *(buffer + encoded) = 0x00 | ((guti->mcc_digit2 & 0x0f) << 4) | (guti->mcc_digit1 & 0x0f);
    MLOG(MDEBUG) << "mcc_digit2 >mcc_digit1 typeofidentity = " << hex << int(*(buffer + encoded));
    encoded++;

    *(buffer + encoded) = 0x00 | ((guti->mnc_digit3 & 0x0f) << 4) | (guti->mcc_digit3 & 0x0f);
    MLOG(MDEBUG) << "mnc_digit3 >mcc_digit3 typeofidentity = " << hex << int(*(buffer + encoded));
    encoded++;
    *(buffer + encoded) = 0x00 | ((guti->mnc_digit2 & 0x0f) << 4) | (guti->mnc_digit1 & 0x0f);
    MLOG(MDEBUG) << "mnc_digit2 >mcc_digit1 typeofidentity = " << hex << int(*(buffer + encoded));
    encoded++;

    *(buffer + encoded) = 0x00 | guti->amfregionid;
    MLOG(MDEBUG) << "amfregionid = " << hex << int(*(buffer + encoded));
    encoded++;
    *(buffer + encoded) = 0x00 | guti->amfsetid;
    MLOG(MDEBUG) << "amfsetid = " << hex << int(*(buffer + encoded));
    encoded++;
    *(buffer + encoded) = 0x00 | ((guti->amfsetid1 & 0x03) << 6) | (guti->amfpointer & 0x3f);
    MLOG(MDEBUG) << "amfsetid1 amfpointer = " << hex << int(*(buffer + encoded));
    encoded++;

    *(buffer + encoded) = 0x00 | guti->tmsi1;
    MLOG(MDEBUG) << "tmsi1 = " << hex << int(*(buffer + encoded));
    encoded++;
    *(buffer + encoded) = 0x00 | guti->tmsi2;
    MLOG(MDEBUG) << "tmsi2 = " << hex << int(*(buffer + encoded));
    encoded++;
    *(buffer + encoded) = 0x00 | guti->tmsi3;
    MLOG(MDEBUG) << "tmsi3 = " << hex << int(*(buffer + encoded));
    encoded++;
    *(buffer + encoded) = 0x00 | guti->tmsi4;
    MLOG(MDEBUG) << "tmsi4 = " << hex << int(*(buffer + encoded));
    encoded++;

    return encoded;
  };

  // TBD
  // Encode ImeiMobileIdentity IE
  int M5GSMobileIdentityMsg::EncodeImeiMobileIdentityMsg(
      ImeiM5GSMobileIdentity* imei, uint8_t* buffer) {
    uint32_t encoded = 0;

    MLOG(MDEBUG) << "EncodeImeiMobileIdentityMsg:";
    *(buffer + encoded) =
      0x00 | ((imei->identity_digit1 & 0xf0) << 4) | ((imei->oddeven & 0x1) << 3) | (imei->typeofidentity & 0x7);
    encoded++;
    *(buffer + encoded) = 0x00 | ((imei->identity_digit2 & 0xf0) << 4) | (imei->identity_digit3 & 0x0f);
    encoded++;

    return encoded;
  };

  // TBD
  // Encode ImsiMobileIdentity IE
  int M5GSMobileIdentityMsg::EncodeImsiMobileIdentityMsg(
      ImsiM5GSMobileIdentity* imsi, uint8_t* buffer) {
    uint32_t encoded = 0;

    MLOG(MDEBUG) << "EncodeImsiMobileIdentityMsg:";

    *(buffer + encoded) =
      0x00 | ((imsi->spare2 & 0x80) << 7) | ((imsi->supiformat & 0x07) << 4) | ((imsi->spare1 & 0x01) << 3) | (imsi->typeofidentity & 0x7);
    encoded++;
    *(buffer + encoded) = 0x00 | ((imsi->mcc_digit2 & 0x0f) << 4) | (imsi->mcc_digit1 & 0x0f);
    encoded++;

    *(buffer + encoded) = 0x00 | ((imsi->mnc_digit3 & 0x0f) << 4) | (imsi->mcc_digit3 & 0x0f);
    encoded++;
    *(buffer + encoded) = 0x00 | ((imsi->mnc_digit2 & 0x0f) << 4) | (imsi->mnc_digit1 & 0x0f);
    encoded++;

    *(buffer + encoded) = 0x00 | ((imsi->routingindicatordigit2 & 0xf0) << 4) | (imsi->routingindicatordigit1 & 0x0f);
    encoded++;
    *(buffer + encoded) = 0x00 | ((imsi->routingindicatordigit3 & 0xf0) << 4) | (imsi->routingindicatordigit4 & 0x0f);
    encoded++;

    *(buffer + encoded) = 0x00 | ((imsi->spare6 & 0x01) << 7) | ((imsi->spare5 & 0x01) << 6) |
      ((imsi->spare4 & 0x01) << 5) | ((imsi->spare3 & 0x01) << 4) | (imsi->protect_schm_id & 0x0f);

    *(buffer + encoded) = imsi->home_nw_id;
    encoded++;

#if 0
    imsi->scheme_output.assign((const char *)(buffer + encoded),  imsi->scheme_output.size());


    MLOG(MDEBUG) << "ielen = " << hex << (unsigned char)imsi->scheme_output.size() ;

    MLOG(MDEBUG) << "contents";
    int i = 0;
    for(i; i < imsi->scheme_output.size(); i++) {
      MLOG(MDEBUG) << (uint8_t)(imsi->scheme_output[i]);
    }
    MLOG(MDEBUG) << endl;
#endif
    return encoded;
  };

  // TBD
  int M5GSMobileIdentityMsg::EncodeTmsiMobileIdentityMsg(
      TmsiM5GSMobileIdentity* tmsi, uint8_t* buffer) {
    uint32_t encoded = 0;

    MLOG(MDEBUG) << "EncodeTmsiMobileIdentityMsg:";
    *(buffer + encoded) =
      0x00 | ((tmsi->spare & 0x0f) << 4) | ((tmsi->oddeven & 0x01) << 3) | (tmsi->typeofidentity & 0x7);
    encoded++;
    *(buffer + encoded) = 0x00 | tmsi->amfsetid;
    encoded++;
    *(buffer + encoded) = 0x00 | ((tmsi->amfsetid1 & 0xc0) << 6);
    *(buffer + encoded) = 0x00 | (tmsi->amfpointer & 0x3f);
    encoded++;

    *(buffer + encoded) = 0x00 | tmsi->m5gtmsi1;
    encoded++;
    *(buffer + encoded) = 0x00 | tmsi->m5gtmsi2;
    encoded++;
    *(buffer + encoded) = 0x00 | tmsi->m5gtmsi3;
    encoded++;
    *(buffer + encoded) = 0x00 | tmsi->m5gtmsi4;
    encoded++;

    return encoded;
  };

  // Encode SuciMobileIdentity IE
  int M5GSMobileIdentityMsg::EncodeSuciMobileIdentityMsg(
      SuciM5GSMobileIdentity* suci, uint8_t* buffer) {
    uint32_t encoded = 0;

    MLOG(MDEBUG) << "EncodeSuciMobileIdentityMsg:";
    *(buffer + encoded) =
      0x00 | ((suci->spare2 & 0x80) << 7) | ((suci->supiformat & 0x07) << 4) | ((suci->spare1 & 0x01) << 3) | (suci->typeofidentity & 0x7);
    encoded++;

    suci->sucinai.assign((const char *)(buffer + encoded), suci->sucinai.size());

    MLOG(MDEBUG) << "ielen = " << hex << (unsigned char)suci->sucinai.size() ;

    MLOG(MDEBUG) << "contents";
    int i = 0;
    for(i; i < suci->sucinai.size(); i++) {
      MLOG(MDEBUG) << hex << int(suci->sucinai[i]);
    }
    MLOG(MDEBUG) << endl;

    return encoded;

  };

  // Encode M5GSMobileIdentity IE
  int M5GSMobileIdentityMsg::EncodeM5GSMobileIdentityMsg(M5GSMobileIdentityMsg *m5gsmobileidentity, uint8_t iei, uint8_t * buffer, uint32_t len)
  {
    uint16_t *lenPtr;
    int encoded_rc = TLV_VALUE_DOESNT_MATCH;
    uint32_t encoded = 0;

    MLOG(MDEBUG) << "EncodeM5GSMobileIdentityMsg:";

    // Checking IEI and pointer

    CHECK_PDU_POINTER_AND_LENGTH_ENCODER(
        buffer, MOBILE_IDENTITY_MIN_LENGTH, len);

    if (iei > 0) {
      CHECK_IEI_ENCODER((unsigned char)iei, m5gsmobileidentity->iei);
      *buffer = iei;
      MLOG(MDEBUG) << "iei" <<  hex << int(*buffer);
      encoded++;
    }

    lenPtr = (uint16_t *)(buffer + encoded);
    encoded += 2;

    if (m5gsmobileidentity->mobileidentity.imsi.typeofidentity == M5GSMobileIdentityMsg_IMSI) {
      MLOG(MDEBUG) << "Type imsi";
      encoded_rc = EncodeImsiMobileIdentityMsg(
          &m5gsmobileidentity->mobileidentity.imsi, buffer + encoded);
    } else if (m5gsmobileidentity->mobileidentity.imei.typeofidentity == M5GSMobileIdentityMsg_IMEI) {
      MLOG(MDEBUG) << "Type imei";
      encoded_rc = EncodeImeiMobileIdentityMsg(
          &m5gsmobileidentity->mobileidentity.imei, buffer + encoded);
    } else if (m5gsmobileidentity->mobileidentity.guti.typeofidentity == M5GSMobileIdentityMsg_GUTI) {
      MLOG(MDEBUG) << "Type guti";
      encoded_rc = EncodeGutiMobileIdentityMsg(
          &m5gsmobileidentity->mobileidentity.guti, buffer + encoded);
    } else if (m5gsmobileidentity->mobileidentity.tmsi.typeofidentity == M5GSMobileIdentityMsg_TMSI) {
      MLOG(MDEBUG) << "Type tmsi";
      encoded_rc = EncodeTmsiMobileIdentityMsg(
          &m5gsmobileidentity->mobileidentity.tmsi, buffer + encoded);
    } else if (m5gsmobileidentity->mobileidentity.suci.typeofidentity == M5GSMobileIdentityMsg_SUCI) {
      MLOG(MDEBUG) << "Type suci";
      encoded_rc = EncodeSuciMobileIdentityMsg(
          &m5gsmobileidentity->mobileidentity.suci, buffer + encoded);
    }

    if (encoded_rc < 0) {
      MLOG(MDEBUG) << "Encode error" << encoded_rc;
      return encoded_rc;
    }

    *lenPtr = htons(encoded + encoded_rc - 2 - ((iei > 0) ? 1 : 0));
    return (encoded + encoded_rc);
  };
}

