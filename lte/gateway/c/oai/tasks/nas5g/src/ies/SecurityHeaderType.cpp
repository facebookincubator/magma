#include <iostream>
#include <sstream>
#include <cstdint>
#include "SecurityHeaderType.h"
#include "CommonDefs.h"

using namespace std;
namespace magma5g {
SecurityHeaderTypeMsg::SecurityHeaderTypeMsg(){};

SecurityHeaderTypeMsg::~SecurityHeaderTypeMsg(){};

// Decode SecurityHeaderType IE
int SecurityHeaderTypeMsg::DecodeSecurityHeaderTypeMsg(
    SecurityHeaderTypeMsg* securityheadertype, uint8_t iei, uint8_t* buffer,
    uint32_t len) {
  int decoded = 0;

  MLOG(MDEBUG) << "   DecodeSecurityHeaderTypeMsg : ";
  securityheadertype->securityhdr = *(buffer) &0xf;
  decoded++;
  MLOG(MDEBUG) << " Security header type = " << dec
               << int(securityheadertype->securityhdr);
  return (decoded);
};

// Encode SecurityHeaderType IE
int SecurityHeaderTypeMsg::EncodeSecurityHeaderTypeMsg(
    SecurityHeaderTypeMsg* securityheadertype, uint8_t iei, uint8_t* buffer,
    uint32_t len) {
  int encoded = 0;

  MLOG(MDEBUG) << " EncodeSecurityHeaderTypeMsg : ";
  *(buffer) = securityheadertype->securityhdr & 0xf;
  encoded++;
  MLOG(MDEBUG) << "Security header type = 0x" << hex << int(*(buffer));
  return (encoded);
};
}  // namespace magma5g
