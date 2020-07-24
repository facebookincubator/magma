/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @flow
 * @format
 */

import * as React from 'react';
import Button from '@fbcnms/ui/components/design-system/Button';
import DialogActions from '@material-ui/core/DialogActions';

type Props = {
  onAbort: () => void,
  onUpload: () => void,
};

const UploadAnywayDialog = (props: Props) => {
  return (
    <div>
      <DialogActions>
        <Button onClick={props.onAbort} skin="regular">
          Cancel
        </Button>
        <Button onClick={props.onUpload}>Upload Anyway</Button>
      </DialogActions>
    </div>
  );
};

export default UploadAnywayDialog;
