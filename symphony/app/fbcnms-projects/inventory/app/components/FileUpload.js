/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import * as React from 'react';
import FormAction from '@fbcnms/ui/components/design-system/Form/FormAction';
import axios from 'axios';
import fbt from 'fbt';
import shortid from 'shortid';
import {FilesUploadContext} from './context/FilesUploadContextProvider';
import {makeStyles} from '@material-ui/styles';
import {useCallback, useContext, useRef} from 'react';

const useStyles = makeStyles(() => ({
  hiddenInput: {
    width: '0px',
    height: '0px',
    opacity: 0,
    overflow: 'hidden',
    position: 'absolute',
    zIndex: -1,
  },
}));

export const FileUploadButton = (props: {
  button: React.Node,
  onFileChanged: (SyntheticEvent<HTMLInputElement>) => void | Promise<void>,
  className?: ?string,
}) => {
  const {button, onFileChanged, className} = props;
  const classes = useStyles();
  const inputRef = useRef();
  const buttonClick = useCallback(() => inputRef?.current?.click(), [inputRef]);
  return (
    <FormAction>
      <input
        className={classes.hiddenInput}
        type="file"
        onChange={onFileChanged}
        ref={inputRef}
        multiple
      />
      <span className={className} onClick={buttonClick}>
        {button}
      </span>
    </FormAction>
  );
};

type Props = {
  button: React.Node,
  className?: ?string,
  onProgress?: (fileId: string, progress: number) => void,
  onFileUploaded: (file: File, key: string) => void,
};

const FileUpload = ({button, onProgress, onFileUploaded, className}: Props) => {
  const uploadContext = useContext(FilesUploadContext);

  const onFileProgress = (fileId, progress) => {
    uploadContext.setFileProgress(fileId, progress);
    onProgress && onProgress(fileId, progress);
  };

  const onFilesChanged = async (e: SyntheticEvent<HTMLInputElement>) => {
    const eventFiles = Array.from(e.currentTarget.files);
    if (!eventFiles || eventFiles.length === 0) {
      return;
    }

    await Promise.all(
      eventFiles.map(async file => {
        const fileId = shortid.generate();
        uploadContext.addFile(fileId, file.name);
        try {
          await uploadFile(fileId, file, onFileUploaded, onFileProgress);
        } catch (e) {
          uploadContext.setFileUploadError(
            fileId,
            fbt(
              'We had a problem uploading this file',
              'Error message describing that we had an error while uploading the file',
            ),
          );
        }
      }),
    );
  };
  return (
    <FileUploadButton
      button={button}
      className={className}
      onFileChanged={async e => await onFilesChanged(e)}
    />
  );
};

export async function uploadFile(
  id: string,
  file: File,
  onUpload: (File, string) => void,
  onProgress?: (fileId: string, progress: number) => void,
) {
  const signingResponse = await axios.get('/store/put', {
    params: {
      contentType: file.type,
    },
  });

  const config = {
    headers: {
      'Content-Type': file.type,
    },
    onUploadProgress: function(progressEvent) {
      const percentCompleted = Math.round(
        (progressEvent.loaded * 100) / progressEvent.total,
      );
      onProgress && onProgress(id, percentCompleted);
    },
  };
  await axios.put(signingResponse.data.URL, file, config);

  onUpload(file, signingResponse.data.key);
}

export default FileUpload;
