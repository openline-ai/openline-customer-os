import React, { FC } from 'react';
import classNames from 'classnames';
import { Editor } from '../../ui-kit/molecules';
import { Controller, useForm } from 'react-hook-form';
import { useCreateOrganizationNote } from '../../../hooks/useNote';
import { editorEmail, editorMode, EditorMode } from '../../../state';
import { EmailFields } from '../../contact/editor/email-fields';
import { useRecoilState, useRecoilValue } from 'recoil';

export enum NoteEditorModes {
  'ADD' = 'ADD',
  'EDIT' = 'EDIT',
}
interface Props {
  mode: NoteEditorModes;
  organizationId: string;
}

const DEFAULT_VALUES = {
  html: '',
  htmlEnhanced: '',
};
export const OrganizationEditor: FC<Props> = ({
  mode = NoteEditorModes.ADD,
  organizationId,
}) => {
  const { handleSubmit, setValue, getValues, control, reset } = useForm({
    defaultValues: DEFAULT_VALUES,
  });
  const [editorModeState, setMode] = useRecoilState(editorMode);
  const {
    handleSubmit: handleSendEmail,
    to,
    respondTo,
  } = useRecoilValue(editorEmail);

  const { onCreateOrganizationNote, saving } = useCreateOrganizationNote({
    organizationId,
  });
  const isEditMode = mode === NoteEditorModes.EDIT;

  const onSubmit = handleSubmit(async (d) => {
    //remove src attribute to not send the file bytes in here
    // identiti header - from session uuid
    const dataToSubmit = {
      appSource: 'Openline',
      html: d?.htmlEnhanced?.replaceAll(/.src(\S*)/g, '') || '',
    };

    editorModeState.mode === EditorMode.Email && handleSendEmail
      ? handleSendEmail(
          dataToSubmit.html.replace(/(<([^>]+)>)/gi, ''),
          () => reset(DEFAULT_VALUES),
          to,
          respondTo,
        )
      : onCreateOrganizationNote(dataToSubmit).then(() =>
          reset(DEFAULT_VALUES),
        );
  });

  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        margin: isEditMode ? '-17px -24px' : 0,
      }}
    >
      <Controller
        name='htmlEnhanced'
        control={control}
        render={({ field }) => (
          <div
            className={classNames({
              'openline-editor-email':
                editorModeState.mode === EditorMode.Email,
            })}
          >
            {editorModeState.mode === EditorMode.Email && <EmailFields />}
            <Editor
              mode={NoteEditorModes.ADD}
              onGetFieldValue={getValues}
              value={field.value}
              saving={saving}
              onSave={onSubmit}
              label={editorModeState.submitButtonLabel}
              onHtmlChanged={(newHtml: string) => setValue('htmlEnhanced', newHtml)}
            />
          </div>
        )}
      />
    </div>
  );
};
