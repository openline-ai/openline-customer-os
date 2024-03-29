import { useState } from 'react';
import { useParams } from 'next/navigation';

import omit from 'lodash/omit';

import { Flex } from '@ui/layout/Flex';
import { useDisclosure } from '@ui/utils';
import { Text } from '@ui/typography/Text';
import { IconButton } from '@ui/form/IconButton';
import { pulseOpacity } from '@ui/utils/keyframes';
import { Collapse } from '@ui/transitions/Collapse';
import { ChevronExpand } from '@ui/media/icons/ChevronExpand';
import { ChevronCollapse } from '@ui/media/icons/ChevronCollapse';

import { Milestones } from './Milestones';
import { checkPlanDone, checkMilestoneDue } from './utils';
import { usePlanMutations } from '../../hooks/usePlanMutations';
import { PlanDatum, MilestoneDatum, NewMilestoneInput } from '../types';
import { useMilestoneMutations } from '../../hooks/useMilestoneMutations';
import {
  PlanMenu,
  PlanDueDate,
  AddMilestoneModal,
  ProgressCompletion,
} from './components';

interface PlanProps {
  plan: PlanDatum;
  isOpen?: boolean;
  onToggle?: (planId: string) => void;
}

export const Plan = ({ plan, isOpen: _isOpen, onToggle }: PlanProps) => {
  const organizationId = useParams()?.id as string;
  const [isHovered, setIsHovered] = useState(false);
  const planMenu = useDisclosure({ id: `${plan.id}-menu` });
  const addMilestoneModal = useDisclosure({ id: `${plan.id}-add-milestone` });
  const [openMilestoneId, setOpenMilestoneId] = useState<string | null>('');

  const { updateOnboardingPlan } = usePlanMutations({
    organizationId,
  });
  const { updateMilestone, addMilestone } = useMilestoneMutations({
    plan,
  });

  const isTemporary = plan.id.startsWith('temp');
  const activeMilestones = plan.milestones
    .filter((m) => !m.retired)
    .sort((a, b) => a.order - b.order);

  const hasMilestones = activeMilestones.length > 0;
  const hasOneMilestone = activeMilestones.length === 1;
  const existingMilestoneNames = activeMilestones.map(
    (milestone) => milestone.name,
  );

  const handleTogglePlan = () => {
    onToggle?.(plan.id);
  };

  const handleToggleMilestone = (id: string) => {
    setOpenMilestoneId((prevId) => (prevId === id ? null : id));
  };

  const handleRemovePlan = () => {
    updateOnboardingPlan.mutate({
      input: {
        id: plan.id,
        retired: true,
        organizationId,
      },
    });
  };

  const handleUpdateMilestone = (milestone: MilestoneDatum) => {
    updateMilestone.mutate({
      input: {
        ...milestone,
        items: milestone.items.map((item) => omit(item, 'id')),
        organizationId,
        organizationPlanId: plan.id,
        updatedAt: new Date().toISOString(),
      },
    });
  };

  const handleRemoveMilestone = (id: string) => {
    const foundMilestone = plan.milestones.find((m) => m.id === id);
    if (!foundMilestone) return;

    updateMilestone.mutate({
      input: {
        ...foundMilestone,
        retired: true,
        organizationId,
        organizationPlanId: plan.id,
        updatedAt: new Date().toISOString(),
        items: foundMilestone.items.map((item) => omit(item, 'id')),
      },
    });
  };

  const handleAddMilestone = (input: NewMilestoneInput) => {
    addMilestone.mutate({
      input: {
        ...input,
        organizationId,
        adhoc: false,
        optional: false,
        organizationPlanId: plan.id,
        createdAt: new Date().toISOString(),
      },
    });
  };

  const isPlanDone = checkPlanDone(plan);
  const nextDueMilestone = plan.milestones.find(checkMilestoneDue);
  const isOpen = (hasOneMilestone && !isPlanDone) || _isOpen;

  return (
    <Flex
      px='3'
      pb='2'
      pt='3'
      w='full'
      bg='gray.50'
      flexDir='column'
      borderRadius='lg'
      border='1px solid'
      borderColor='gray.200'
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => !isOpen && setIsHovered(false)}
      animation={
        isTemporary ? `${pulseOpacity} 0.7s alternate ease-in-out` : undefined
      }
    >
      <Flex mx='1' align='center' flexDir='column' alignItems='flex-start'>
        <Flex align='center' justify='space-between' w='full'>
          <Text fontSize='sm' fontWeight='semibold' noOfLines={1}>
            {plan.name}
          </Text>

          <Flex
            align='center'
            opacity={isHovered || planMenu.isOpen ? '1' : '0'}
            transition='opacity 0.2s ease-out'
          >
            {((hasMilestones && !hasOneMilestone) || isPlanDone) && (
              <IconButton
                size='xs'
                variant='ghost'
                color='gray.500'
                aria-label='Toggle plan'
                onClick={handleTogglePlan}
                icon={
                  isOpen ? (
                    <ChevronCollapse color='gray.400' />
                  ) : (
                    <ChevronExpand color='gray.400' />
                  )
                }
              />
            )}
            <PlanMenu
              id={plan.masterPlanId}
              isOpen={planMenu.isOpen}
              onOpen={planMenu.onOpen}
              onClose={planMenu.onClose}
              onRemovePlan={handleRemovePlan}
              onAddMilestone={addMilestoneModal.onOpen}
            />
          </Flex>
        </Flex>

        <Flex gap='1'>
          {hasMilestones ? (
            <>
              {nextDueMilestone && (
                <Text fontSize='sm' color='gray.500' fontWeight='semibold'>
                  {nextDueMilestone?.name}
                </Text>
              )}
              <PlanDueDate
                isDone={isPlanDone}
                status={nextDueMilestone?.statusDetails?.status}
                value={
                  isPlanDone
                    ? plan.statusDetails.updatedAt
                    : nextDueMilestone?.dueDate
                }
              />
            </>
          ) : (
            <Text fontSize='sm' color='gray.500'>
              No milestones added yet
            </Text>
          )}
        </Flex>
      </Flex>

      <Collapse in={isOpen}>
        <Milestones
          milestones={activeMilestones}
          openMilestoneId={openMilestoneId}
          onSyncMilestone={handleUpdateMilestone}
          onRemoveMilestone={handleRemoveMilestone}
          onToggleMilestone={handleToggleMilestone}
        />
      </Collapse>

      <Collapse in={!isOpen}>
        <ProgressCompletion plan={plan} />
      </Collapse>

      <AddMilestoneModal
        masterPlanId={plan.masterPlanId}
        isOpen={addMilestoneModal.isOpen}
        onClose={addMilestoneModal.onClose}
        onAddMilestone={handleAddMilestone}
        existingMilestoneNames={existingMilestoneNames}
      />
    </Flex>
  );
};
