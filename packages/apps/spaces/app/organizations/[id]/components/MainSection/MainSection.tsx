'use client';
import { CardHeader, Card, CardBody } from '@ui/presentation/Card';
import { Heading } from '@ui/typography/Heading';
import React from 'react';

export const MainSection = ({ children }: { children?: React.ReactNode }) => {
  return (
    <Card
      flex='3'
      h='calc(100vh - 1rem)'
      bg='#FCFCFC'
      borderRadius='2xl'
      flexDirection='column'
      boxShadow='none'
      position='relative'
      background='gray.25'
      minWidth={609}
      padding={0}
    >
      <CardHeader px={6} pb={2}>
        <Heading as='h1' fontSize='lg' color='gray.700'>
          Timeline
        </Heading>
      </CardHeader>
      <CardBody pr={0} pt={0} p={0} position='unset'>
        {children}
      </CardBody>
    </Card>
  );
};
