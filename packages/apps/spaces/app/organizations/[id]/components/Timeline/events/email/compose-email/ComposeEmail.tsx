'use client';
import React, { FC, useRef } from 'react';

import { ModeChangeButtons } from '@organization/components/Timeline/events/email/compose-email/EmailResponseModeChangeButtons';
import { Box } from '@ui/layout/Box';
import { ParticipantsSelectGroup } from '@organization/components/Timeline/events/email/compose-email/ParticipantsSelectGroup';
import { RichTextEditor } from '@ui/form/RichTextEditor/RichTextEditor';
import { BasicEditorMenu } from '@ui/form/RichTextEditor/menu/BasicEditorMenu';
import {
  BasicEditorExtentions,
  RemirrorProps,
} from '@ui/form/RichTextEditor/types';

interface ComposeEmail {
  onModeChange?: (status: 'reply' | 'reply-all' | 'forward') => void;
  onSubmit: () => void;
  formId: string;
  modal: boolean;
  isSending: boolean;
  to: Array<{ label: string; value: string }>;
  cc: Array<{ label: string; value: string }>;
  bcc: Array<{ label: string; value: string }>;
  remirrorProps: RemirrorProps<BasicEditorExtentions>;
}

export const ComposeEmail: FC<ComposeEmail> = ({
  onModeChange,
  formId,
  modal,
  isSending,
  onSubmit,
  to,
  cc,
  bcc,
  remirrorProps,
}) => {
  const myRef = useRef<HTMLDivElement>(null);
  const height =
    modal && (myRef?.current?.getBoundingClientRect()?.height || 0) + 96;

  return (
    <Box
      borderTop={modal ? '1px dashed var(--gray-200, #EAECF0)' : 'none'}
      background={modal ? '#F8F9FC' : 'white'}
      borderRadius={modal ? 0 : 'lg'}
      borderBottomRadius='2xl'
      as='form'
      p={5}
      overflow='visible'
      maxHeight={modal ? '50vh' : 'auto'}
      pt={1}
      onSubmit={(e) => {
        e.preventDefault();
      }}
    >
      {!!onModeChange && (
        <div style={{ position: 'relative' }}>
          <ModeChangeButtons handleModeChange={onModeChange} />
        </div>
      )}
      <Box ref={myRef}>
        <ParticipantsSelectGroup
          to={to}
          cc={cc}
          bcc={bcc}
          modal={modal}
          formId={formId}
        />
      </Box>

      {/*<FormAutoresizeTextarea*/}
      {/*  placeholder='Write something here...'*/}
      {/*  size='md'*/}
      {/*  formId={formId}*/}
      {/*  name='content'*/}
      {/*  mb={3}*/}
      {/*  transform={!modal ? 'translateY(-16px)' : undefined}*/}
      {/*  resize='none'*/}
      {/*  borderBottom='none'*/}
      {/*  outline='none'*/}
      {/*  borderBottomWidth={0}*/}
      {/*  minHeight={modal ? '100px' : '30px'}*/}
      {/*  maxHeight={modal ? `calc(50vh - ${height}px) !important` : 'auto'}*/}
      {/*  height={modal ? `calc(50vh - ${height}px) !important` : 'auto'}*/}
      {/*  position='initial'*/}
      {/*  overflowY='auto'*/}
      {/*  _focusVisible={{*/}
      {/*    boxShadow: 'none',*/}
      {/*  }}*/}
      {/*/>*/}

      <Box
        // minHeight={modal ? '100px' : '30px'}
        maxHeight={modal ? `calc(50vh - ${height}px) !important` : 'auto'}
        // style={{'--remirror-editor-max-height': modal ? `calc(50vh - ${height}px) !important` : 'auto'}}
        // height={modal ? `calc(50vh - ${height}px) !important` : 'auto'}
        w='full'
      >
        <RichTextEditor {...remirrorProps} formId={formId} name='content'>
          <BasicEditorMenu isSending={isSending} onSubmit={onSubmit} />
        </RichTextEditor>
      </Box>

      {/*<Flex*/}
      {/*  justifyContent='flex-end'*/}
      {/*  direction='row'*/}
      {/*  flex={1}*/}
      {/*  mt='lg'*/}
      {/*  width='100%'*/}
      {/*>*/}

      {/*</Flex>*/}
      {/*{isUploadAreaOpen && (*/}
      {/*  <FileUpload*/}
      {/*    files={files}*/}
      {/*    onBeginFileUpload={(fileKey: string) => {*/}
      {/*      setFiles((prevFiles: any) => [*/}
      {/*        ...prevFiles,*/}
      {/*        {*/}
      {/*          key: fileKey,*/}
      {/*          uploaded: false,*/}
      {/*        },*/}
      {/*      ]);*/}
      {/*    }}*/}
      {/*    onFileUpload={(newFile: any) => {*/}
      {/*      setFiles((prevFiles: any) => {*/}
      {/*        return prevFiles.map((file: any) => {*/}
      {/*          if (file.key === newFile.key) {*/}
      {/*            file = {*/}
      {/*              id: newFile.id,*/}
      {/*              key: newFile.key,*/}
      {/*              name: newFile.name,*/}
      {/*              extension: newFile.extension,*/}
      {/*              uploaded: true,*/}
      {/*            };*/}
      {/*          }*/}
      {/*          return file;*/}
      {/*        });*/}
      {/*      });*/}
      {/*    }}*/}
      {/*    onFileUploadError={(fileKey: any) => {*/}
      {/*      setFiles((prevFiles: any) => {*/}
      {/*        // TODO do not remove the file from the list*/}
      {/*        // show the error instead for that particular file*/}
      {/*        return prevFiles.filter((file: any) => file.key !== fileKey);*/}
      {/*      });*/}
      {/*    }}*/}
      {/*    onFileRemove={(fileId: any) => {*/}
      {/*      setFiles((prevFiles: any) => {*/}
      {/*        return prevFiles.filter((file: any) => file.id !== fileId);*/}
      {/*      });*/}
      {/*    }}*/}
      {/*  />*/}
      {/*)}*/}
    </Box>
  );
};
