'use client';
import { OptionsOrGroups } from 'react-select';
import { useField } from 'react-inverted-form';
import React, { FC, useRef, useState, useEffect, KeyboardEvent } from 'react';

import { AnimatePresence } from 'framer-motion';
import { GroupBase, OptionProps } from 'chakra-react-select';

import { Flex } from '@ui/layout/Flex';
import { Text } from '@ui/typography/Text';
import { chakraComponents } from '@ui/form/SyncSelect';
import { MultiCreatableSelect } from '@ui/form/MultiCreatableSelect';

import { TagButton } from './TagButton';
import { tagsSelectStyles } from './tagSelectStyles';
import { useTagButtonSlideAnimation } from './useTagButtonSlideAnimation';

interface EmailParticipantSelect {
  name: string;
  formId: string;
  tags?: Array<{ value: string; label: string }>;
}

interface Tag {
  label: string;
  value: string;
}
export const suggestedTags = [
  'meeting',
  'call',
  'voicemail',
  'email',
  'text-message',
];

export const TagsSelect: FC<EmailParticipantSelect> = ({
  formId,
  name,
  tags = [],
}) => {
  const { getInputProps } = useField(name, formId);
  const { onChange, value: selectedTags, onBlur } = getInputProps();
  const [isMenuOpen, setMenuOpen] = useState(false);
  const [focusedOption, setFocusedOption] = useState<Tag | null>(null);
  const [inputVal, setInputVal] = useState('');
  const scope = useTagButtonSlideAnimation(!!selectedTags?.length);

  const getFilteredSuggestions = (
    filterString: string,
    callback: (options: OptionsOrGroups<unknown, GroupBase<unknown>>) => void,
  ) => {
    if (!filterString.slice(1).length) {
      callback(tags);

      return;
    }

    const options: OptionsOrGroups<unknown, GroupBase<unknown>> = tags.filter(
      (e) =>
        e.label.toLowerCase().includes(filterString.slice(1)?.toLowerCase()),
    );

    callback(options);
  };
  const handleInputChange = (d: string) => {
    setInputVal(d);
    if (d.length === 1 && d.startsWith('#')) {
      setMenuOpen(true);
    }
    if (!d.length || !d.startsWith('#')) {
      setMenuOpen(false);
    }
  };

  // this function is needed as tags are selected on 'Space' & 'Enter'
  const handleKeyDown = (event: KeyboardEvent) => {
    if (event.code === 'Backspace') {
      if (inputVal.length) {
        return;
      }
      event.preventDefault();
      const newSelected = [...selectedTags].slice(0, selectedTags.length - 1);
      onChange(newSelected);
    }
    if (event.code === 'Space' || event.code === 'Enter') {
      event.preventDefault();
      if (!isMenuOpen) return;

      if (focusedOption) {
        onChange([...selectedTags, focusedOption]);
        setMenuOpen(false);
        setFocusedOption(null);
        setInputVal('');
      }
    }
  };

  // FIXME - move this to outer scope
  const Option = (props: OptionProps<{ label: string; value: string }>) => {
    const Or = useRef(null);

    useEffect(() => {
      if (props.isFocused) {
        setFocusedOption(props.data);
      }
    }, [props.isFocused, props.data.label]);

    return (
      <div ref={Or}>
        <chakraComponents.Option {...props} key={props.data.label}>
          {props.data.label || props.data.value}
        </chakraComponents.Option>
      </div>
    );
  };

  return (
    <>
      <AnimatePresence initial={false}>
        <Flex alignItems='baseline' ref={scope}>
          {!selectedTags?.length && (
            <>
              <Text color='gray.500' mr={2} whiteSpace='nowrap'>
                Suggested tags:
              </Text>

              {suggestedTags?.map((tag) => (
                <TagButton
                  key={`tag-select-${tag}`}
                  onTagSet={() =>
                    onChange([
                      {
                        label: tag,
                        value:
                          tags?.find((e) => suggestedTags.includes(e.label))
                            ?.value || tag,
                      },
                    ])
                  }
                  tag={tag}
                />
              ))}
            </>
          )}
          {!!selectedTags?.length && (
            <MultiCreatableSelect
              Option={Option}
              name={name}
              formId={formId}
              placeholder=''
              backspaceRemovesValue
              onKeyDown={handleKeyDown}
              onChange={onChange}
              noOptionsMessage={() => null}
              loadOptions={(inputValue: string, callback) => {
                getFilteredSuggestions(inputValue, callback);
              }}
              formatCreateLabel={(input) => {
                if (input?.startsWith('#')) {
                  return `${input.slice(1)}`;
                }

                return input;
              }}
              onBlur={() => onBlur(selectedTags)}
              onMenuClose={() => setFocusedOption(null)}
              value={selectedTags}
              inputValue={inputVal}
              onInputChange={handleInputChange}
              menuIsOpen={isMenuOpen}
              menuPlacement='top'
              defaultOptions={tags}
              hideSelectedOptions
              isValidNewOption={(input) => input.startsWith('#')}
              getOptionLabel={(d) => {
                if (d.label?.startsWith('#')) {
                  return `${d.label.slice(1)}`;
                }

                return `${d.label}`;
              }}
              menuShouldBlockScroll
              onCreateOption={(input) => {
                if (input?.startsWith('#')) {
                  return {
                    value: input.slice(1),
                    label: input.slice(1),
                  };
                }

                return {
                  value: input,
                  label: input,
                };
              }}
              getNewOptionData={(input) => {
                if (input?.startsWith('#')) {
                  return {
                    value: input.slice(1),
                    label: input.slice(1),
                  };
                }

                return {
                  value: input,
                  label: input,
                };
              }}
              // @ts-expect-error remove this in favour of chakraStyles
              customStyles={tagsSelectStyles}
            />
          )}
        </Flex>
      </AnimatePresence>
    </>
  );
};
