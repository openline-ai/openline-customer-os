export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
export type MakeEmpty<
  T extends { [key: string]: unknown },
  K extends keyof T,
> = { [_ in K]?: never };
export type Incremental<T> =
  | T
  | {
      [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never;
    };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string };
  String: { input: string; output: string };
  Boolean: { input: boolean; output: boolean };
  Int: { input: number; output: number };
  Float: { input: number; output: number };
  Any: { input: any; output: any };
  Int64: { input: any; output: any };
  Time: { input: any; output: any };
};

export type Action = {
  __typename?: 'Action';
  actionType: ActionType;
  appSource: Scalars['String']['output'];
  content?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  id: Scalars['ID']['output'];
  metadata?: Maybe<Scalars['String']['output']>;
  source: DataSource;
};

export type ActionItem = {
  __typename?: 'ActionItem';
  appSource: Scalars['String']['output'];
  content: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  source: DataSource;
};

export enum ActionType {
  ContractRenewed = 'CONTRACT_RENEWED',
  ContractStatusUpdated = 'CONTRACT_STATUS_UPDATED',
  Created = 'CREATED',
  OnboardingStatusChanged = 'ONBOARDING_STATUS_CHANGED',
  RenewalForecastUpdated = 'RENEWAL_FORECAST_UPDATED',
  RenewalLikelihoodUpdated = 'RENEWAL_LIKELIHOOD_UPDATED',
  ServiceLineItemBilledTypeOnceCreated = 'SERVICE_LINE_ITEM_BILLED_TYPE_ONCE_CREATED',
  ServiceLineItemBilledTypeRecurringCreated = 'SERVICE_LINE_ITEM_BILLED_TYPE_RECURRING_CREATED',
  ServiceLineItemBilledTypeUpdated = 'SERVICE_LINE_ITEM_BILLED_TYPE_UPDATED',
  ServiceLineItemBilledTypeUsageCreated = 'SERVICE_LINE_ITEM_BILLED_TYPE_USAGE_CREATED',
  ServiceLineItemPriceUpdated = 'SERVICE_LINE_ITEM_PRICE_UPDATED',
  ServiceLineItemQuantityUpdated = 'SERVICE_LINE_ITEM_QUANTITY_UPDATED',
  ServiceLineItemRemoved = 'SERVICE_LINE_ITEM_REMOVED',
}

export type Analysis = Node & {
  __typename?: 'Analysis';
  analysisType?: Maybe<Scalars['String']['output']>;
  appSource: Scalars['String']['output'];
  content?: Maybe<Scalars['String']['output']>;
  contentType?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  describes: Array<DescriptionNode>;
  id: Scalars['ID']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
};

export type AnalysisDescriptionInput = {
  interactionEventId?: InputMaybe<Scalars['ID']['input']>;
  interactionSessionId?: InputMaybe<Scalars['ID']['input']>;
  meetingId?: InputMaybe<Scalars['ID']['input']>;
};

export type AnalysisInput = {
  analysisType?: InputMaybe<Scalars['String']['input']>;
  appSource: Scalars['String']['input'];
  content?: InputMaybe<Scalars['String']['input']>;
  contentType?: InputMaybe<Scalars['String']['input']>;
  describes: Array<AnalysisDescriptionInput>;
};

export type Attachment = Node & {
  __typename?: 'Attachment';
  appSource: Scalars['String']['output'];
  basePath: Scalars['String']['output'];
  cdnUrl: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  fileName: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  mimeType: Scalars['String']['output'];
  size: Scalars['Int64']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
};

export type AttachmentInput = {
  appSource: Scalars['String']['input'];
  basePath: Scalars['String']['input'];
  cdnUrl: Scalars['String']['input'];
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  fileName: Scalars['String']['input'];
  id?: InputMaybe<Scalars['ID']['input']>;
  mimeType: Scalars['String']['input'];
  size: Scalars['Int64']['input'];
};

export type BankAccount = MetadataInterface & {
  __typename?: 'BankAccount';
  accountNumber?: Maybe<Scalars['String']['output']>;
  allowInternational: Scalars['Boolean']['output'];
  bankName?: Maybe<Scalars['String']['output']>;
  bankTransferEnabled: Scalars['Boolean']['output'];
  bic?: Maybe<Scalars['String']['output']>;
  currency?: Maybe<Currency>;
  iban?: Maybe<Scalars['String']['output']>;
  metadata: Metadata;
  otherDetails?: Maybe<Scalars['String']['output']>;
  routingNumber?: Maybe<Scalars['String']['output']>;
  sortCode?: Maybe<Scalars['String']['output']>;
};

export type BankAccountCreateInput = {
  accountNumber?: InputMaybe<Scalars['String']['input']>;
  allowInternational?: InputMaybe<Scalars['Boolean']['input']>;
  bankName?: InputMaybe<Scalars['String']['input']>;
  bankTransferEnabled?: InputMaybe<Scalars['Boolean']['input']>;
  bic?: InputMaybe<Scalars['String']['input']>;
  currency?: InputMaybe<Currency>;
  iban?: InputMaybe<Scalars['String']['input']>;
  otherDetails?: InputMaybe<Scalars['String']['input']>;
  routingNumber?: InputMaybe<Scalars['String']['input']>;
  sortCode?: InputMaybe<Scalars['String']['input']>;
};

export type BankAccountUpdateInput = {
  accountNumber?: InputMaybe<Scalars['String']['input']>;
  allowInternational?: InputMaybe<Scalars['Boolean']['input']>;
  bankName?: InputMaybe<Scalars['String']['input']>;
  bankTransferEnabled?: InputMaybe<Scalars['Boolean']['input']>;
  bic?: InputMaybe<Scalars['String']['input']>;
  currency?: InputMaybe<Currency>;
  iban?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['ID']['input'];
  otherDetails?: InputMaybe<Scalars['String']['input']>;
  routingNumber?: InputMaybe<Scalars['String']['input']>;
  sortCode?: InputMaybe<Scalars['String']['input']>;
};

export enum BilledType {
  Annually = 'ANNUALLY',
  Monthly = 'MONTHLY',
  None = 'NONE',
  Once = 'ONCE',
  Quarterly = 'QUARTERLY',
  Usage = 'USAGE',
}

export type BillingDetails = {
  __typename?: 'BillingDetails';
  addressLine1?: Maybe<Scalars['String']['output']>;
  addressLine2?: Maybe<Scalars['String']['output']>;
  billingCycle?: Maybe<ContractBillingCycle>;
  billingEmail?: Maybe<Scalars['String']['output']>;
  canPayWithBankTransfer?: Maybe<Scalars['Boolean']['output']>;
  canPayWithCard?: Maybe<Scalars['Boolean']['output']>;
  canPayWithDirectDebit?: Maybe<Scalars['Boolean']['output']>;
  country?: Maybe<Scalars['String']['output']>;
  invoiceNote?: Maybe<Scalars['String']['output']>;
  invoicingStarted?: Maybe<Scalars['Time']['output']>;
  locality?: Maybe<Scalars['String']['output']>;
  nextInvoicing?: Maybe<Scalars['Time']['output']>;
  organizationLegalName?: Maybe<Scalars['String']['output']>;
  payAutomatically?: Maybe<Scalars['Boolean']['output']>;
  payOnline?: Maybe<Scalars['Boolean']['output']>;
  postalCode?: Maybe<Scalars['String']['output']>;
  region?: Maybe<Scalars['String']['output']>;
};

export type BillingDetailsInput = {
  addressLine1?: InputMaybe<Scalars['String']['input']>;
  addressLine2?: InputMaybe<Scalars['String']['input']>;
  billingCycle?: InputMaybe<ContractBillingCycle>;
  billingEmail?: InputMaybe<Scalars['String']['input']>;
  canPayWithBankTransfer?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithCard?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithDirectDebit?: InputMaybe<Scalars['Boolean']['input']>;
  country?: InputMaybe<Scalars['String']['input']>;
  invoiceNote?: InputMaybe<Scalars['String']['input']>;
  invoicingStarted?: InputMaybe<Scalars['Time']['input']>;
  locality?: InputMaybe<Scalars['String']['input']>;
  organizationLegalName?: InputMaybe<Scalars['String']['input']>;
  payAutomatically?: InputMaybe<Scalars['Boolean']['input']>;
  payOnline?: InputMaybe<Scalars['Boolean']['input']>;
  postalCode?: InputMaybe<Scalars['String']['input']>;
  region?: InputMaybe<Scalars['String']['input']>;
};

export type BillingProfile = Node &
  SourceFields & {
    __typename?: 'BillingProfile';
    appSource: Scalars['String']['output'];
    createdAt: Scalars['Time']['output'];
    id: Scalars['ID']['output'];
    legalName: Scalars['String']['output'];
    source: DataSource;
    sourceOfTruth: DataSource;
    taxId: Scalars['String']['output'];
    updatedAt: Scalars['Time']['output'];
  };

export type BillingProfileInput = {
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  legalName?: InputMaybe<Scalars['String']['input']>;
  organizationId: Scalars['ID']['input'];
  taxId?: InputMaybe<Scalars['String']['input']>;
};

export type BillingProfileLinkEmailInput = {
  billingProfileId: Scalars['ID']['input'];
  emailId: Scalars['ID']['input'];
  organizationId: Scalars['ID']['input'];
  primary?: InputMaybe<Scalars['Boolean']['input']>;
};

export type BillingProfileLinkLocationInput = {
  billingProfileId: Scalars['ID']['input'];
  locationId: Scalars['ID']['input'];
  organizationId: Scalars['ID']['input'];
};

export type BillingProfileUpdateInput = {
  billingProfileId: Scalars['ID']['input'];
  legalName?: InputMaybe<Scalars['String']['input']>;
  organizationId: Scalars['ID']['input'];
  taxId?: InputMaybe<Scalars['String']['input']>;
  updatedAt?: InputMaybe<Scalars['Time']['input']>;
};

/**
 * Describes the relationship a Contact has with a Organization.
 * **A `return` object**
 */
export type Calendar = {
  __typename?: 'Calendar';
  appSource: Scalars['String']['output'];
  calType: CalendarType;
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  link?: Maybe<Scalars['String']['output']>;
  primary: Scalars['Boolean']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
  updatedAt: Scalars['Time']['output'];
};

export enum CalendarType {
  Calcom = 'CALCOM',
  Google = 'GOOGLE',
}

export type ColumnDef = Node & {
  __typename?: 'ColumnDef';
  columnType?: Maybe<ColumnType>;
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  id: Scalars['ID']['output'];
  isDefaultSort?: Maybe<Scalars['Boolean']['output']>;
  isFilterable?: Maybe<Scalars['Boolean']['output']>;
  isSortable?: Maybe<Scalars['Boolean']['output']>;
  isVisible?: Maybe<Scalars['Boolean']['output']>;
  type?: Maybe<ViewType>;
  updatedAt: Scalars['Time']['output'];
};

export type ColumnType = Node & {
  __typename?: 'ColumnType';
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  id: Scalars['ID']['output'];
  name?: Maybe<Scalars['String']['output']>;
  updatedAt: Scalars['Time']['output'];
  viewTypeId?: Maybe<Scalars['String']['output']>;
};

export type Comment = {
  __typename?: 'Comment';
  appSource: Scalars['String']['output'];
  content?: Maybe<Scalars['String']['output']>;
  contentType?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  externalLinks: Array<ExternalSystem>;
  id: Scalars['ID']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
  updatedAt: Scalars['Time']['output'];
};

export enum ComparisonOperator {
  Between = 'BETWEEN',
  Contains = 'CONTAINS',
  Eq = 'EQ',
  Gte = 'GTE',
  In = 'IN',
  IsEmpty = 'IS_EMPTY',
  IsNull = 'IS_NULL',
  Lte = 'LTE',
  StartsWith = 'STARTS_WITH',
}

/**
 * A contact represents an individual in customerOS.
 * **A `response` object.**
 */
export type Contact = ExtensibleEntity &
  Node & {
    __typename?: 'Contact';
    appSource?: Maybe<Scalars['String']['output']>;
    /**
     * An ISO8601 timestamp recording when the contact was created in customerOS.
     * **Required**
     */
    createdAt: Scalars['Time']['output'];
    /**
     * User defined metadata appended to the contact record in customerOS.
     * **Required.  If no values it returns an empty array.**
     */
    customFields: Array<CustomField>;
    description?: Maybe<Scalars['String']['output']>;
    /**
     * All email addresses associated with a contact in customerOS.
     * **Required.  If no values it returns an empty array.**
     */
    emails: Array<Email>;
    fieldSets: Array<FieldSet>;
    /** The first name of the contact in customerOS. */
    firstName?: Maybe<Scalars['String']['output']>;
    /**
     * The unique ID associated with the contact in customerOS.
     * **Required**
     */
    id: Scalars['ID']['output'];
    /**
     * `organizationName` and `jobTitle` of the contact if it has been associated with an organization.
     * **Required.  If no values it returns an empty array.**
     */
    jobRoles: Array<JobRole>;
    /** @deprecated Use `tags` instead */
    label?: Maybe<Scalars['String']['output']>;
    /** The last name of the contact in customerOS. */
    lastName?: Maybe<Scalars['String']['output']>;
    /**
     * All locations associated with a contact in customerOS.
     * **Required.  If no values it returns an empty array.**
     */
    locations: Array<Location>;
    /** The name of the contact in customerOS, alternative for firstName + lastName. */
    name?: Maybe<Scalars['String']['output']>;
    /** Contact notes */
    notes: NotePage;
    notesByTime: Array<Note>;
    organizations: OrganizationPage;
    /** Contact owner (user) */
    owner?: Maybe<User>;
    /**
     * All phone numbers associated with a contact in customerOS.
     * **Required.  If no values it returns an empty array.**
     */
    phoneNumbers: Array<PhoneNumber>;
    prefix?: Maybe<Scalars['String']['output']>;
    profilePhotoUrl?: Maybe<Scalars['String']['output']>;
    socials: Array<Social>;
    source: DataSource;
    sourceOfTruth: DataSource;
    tags?: Maybe<Array<Tag>>;
    /** Template of the contact in customerOS. */
    template?: Maybe<EntityTemplate>;
    timelineEvents: Array<TimelineEvent>;
    timelineEventsTotalCount: Scalars['Int64']['output'];
    timezone?: Maybe<Scalars['String']['output']>;
    /**
     * The title associate with the contact in customerOS.
     * @deprecated Use `prefix` instead
     */
    title?: Maybe<Scalars['String']['output']>;
    updatedAt: Scalars['Time']['output'];
  };

/**
 * A contact represents an individual in customerOS.
 * **A `response` object.**
 */
export type ContactNotesArgs = {
  pagination?: InputMaybe<Pagination>;
};

/**
 * A contact represents an individual in customerOS.
 * **A `response` object.**
 */
export type ContactNotesByTimeArgs = {
  pagination?: InputMaybe<TimeRange>;
};

/**
 * A contact represents an individual in customerOS.
 * **A `response` object.**
 */
export type ContactOrganizationsArgs = {
  pagination?: InputMaybe<Pagination>;
  sort?: InputMaybe<Array<SortBy>>;
  where?: InputMaybe<Filter>;
};

/**
 * A contact represents an individual in customerOS.
 * **A `response` object.**
 */
export type ContactTimelineEventsArgs = {
  from?: InputMaybe<Scalars['Time']['input']>;
  size: Scalars['Int']['input'];
  timelineEventTypes?: InputMaybe<Array<TimelineEventType>>;
};

/**
 * A contact represents an individual in customerOS.
 * **A `response` object.**
 */
export type ContactTimelineEventsTotalCountArgs = {
  timelineEventTypes?: InputMaybe<Array<TimelineEventType>>;
};

/**
 * Create an individual in customerOS.
 * **A `create` object.**
 */
export type ContactInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  /** An ISO8601 timestamp recording when the contact was created in customerOS. */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  /**
   * User defined metadata appended to contact.
   * **Required.**
   */
  customFields?: InputMaybe<Array<CustomFieldInput>>;
  description?: InputMaybe<Scalars['String']['input']>;
  /** An email addresses associated with the contact. */
  email?: InputMaybe<EmailInput>;
  externalReference?: InputMaybe<ExternalSystemReferenceInput>;
  fieldSets?: InputMaybe<Array<FieldSetInput>>;
  /** The first name of the contact. */
  firstName?: InputMaybe<Scalars['String']['input']>;
  /** The last name of the contact. */
  lastName?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  /** Id of the contact owner (user) */
  ownerId?: InputMaybe<Scalars['ID']['input']>;
  /** A phone number associated with the contact. */
  phoneNumber?: InputMaybe<PhoneNumberInput>;
  /** The prefix of the contact. */
  prefix?: InputMaybe<Scalars['String']['input']>;
  profilePhotoUrl?: InputMaybe<Scalars['String']['input']>;
  /** The unique ID associated with the template of the contact in customerOS. */
  templateId?: InputMaybe<Scalars['ID']['input']>;
  timezone?: InputMaybe<Scalars['String']['input']>;
};

export type ContactOrganizationInput = {
  contactId: Scalars['ID']['input'];
  organizationId: Scalars['ID']['input'];
};

export type ContactParticipant = {
  __typename?: 'ContactParticipant';
  contactParticipant: Contact;
  type?: Maybe<Scalars['String']['output']>;
};

export type ContactTagInput = {
  contactId: Scalars['ID']['input'];
  tagId: Scalars['ID']['input'];
};

/**
 * Updates data fields associated with an existing customer record in customerOS.
 * **An `update` object.**
 */
export type ContactUpdateInput = {
  description?: InputMaybe<Scalars['String']['input']>;
  firstName?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['ID']['input'];
  lastName?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  patch?: InputMaybe<Scalars['Boolean']['input']>;
  prefix?: InputMaybe<Scalars['String']['input']>;
  profilePhotoUrl?: InputMaybe<Scalars['String']['input']>;
  timezone?: InputMaybe<Scalars['String']['input']>;
};

/**
 * Specifies how many pages of contact information has been returned in the query response.
 * **A `response` object.**
 */
export type ContactsPage = Pages & {
  __typename?: 'ContactsPage';
  /**
   * A contact entity in customerOS.
   * **Required.  If no values it returns an empty array.**
   */
  content: Array<Contact>;
  /**
   * Total number of elements in the query response.
   * **Required.**
   */
  totalElements: Scalars['Int64']['output'];
  /**
   * Total number of pages in the query response.
   * **Required.**
   */
  totalPages: Scalars['Int']['output'];
};

export type Contract = MetadataInterface & {
  __typename?: 'Contract';
  /** @deprecated Use billingDetails instead. */
  addressLine1?: Maybe<Scalars['String']['output']>;
  /** @deprecated Use billingDetails instead. */
  addressLine2?: Maybe<Scalars['String']['output']>;
  /** @deprecated Use metadata instead. */
  appSource: Scalars['String']['output'];
  /** @deprecated Use billingDetails instead. */
  billingCycle?: Maybe<ContractBillingCycle>;
  billingDetails?: Maybe<BillingDetails>;
  billingEnabled: Scalars['Boolean']['output'];
  committedPeriods?: Maybe<Scalars['Int64']['output']>;
  contractEnded?: Maybe<Scalars['Time']['output']>;
  contractLineItems?: Maybe<Array<ServiceLineItem>>;
  contractName: Scalars['String']['output'];
  contractRenewalCycle: ContractRenewalCycle;
  contractSigned?: Maybe<Scalars['Time']['output']>;
  contractStatus: ContractStatus;
  contractUrl?: Maybe<Scalars['String']['output']>;
  /** @deprecated Use billingDetails instead. */
  country?: Maybe<Scalars['String']['output']>;
  /** @deprecated Use metadata instead. */
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  currency?: Maybe<Currency>;
  /** @deprecated Use contractEnded instead. */
  endedAt?: Maybe<Scalars['Time']['output']>;
  externalLinks: Array<ExternalSystem>;
  /** @deprecated Use metadata instead. */
  id: Scalars['ID']['output'];
  /** @deprecated Use billingDetails instead. */
  invoiceEmail?: Maybe<Scalars['String']['output']>;
  /** @deprecated Use billingDetails instead. */
  invoiceNote?: Maybe<Scalars['String']['output']>;
  /** @deprecated Use billingDetails instead. */
  invoicingStartDate?: Maybe<Scalars['Time']['output']>;
  /** @deprecated Use billingDetails instead. */
  locality?: Maybe<Scalars['String']['output']>;
  metadata: Metadata;
  /** @deprecated Use contractName instead. */
  name: Scalars['String']['output'];
  opportunities?: Maybe<Array<Opportunity>>;
  /** @deprecated Use billingDetails instead. */
  organizationLegalName?: Maybe<Scalars['String']['output']>;
  owner?: Maybe<User>;
  /** @deprecated Use contractRenewalCycle instead. */
  renewalCycle: ContractRenewalCycle;
  /** @deprecated Use committedPeriods instead. */
  renewalPeriods?: Maybe<Scalars['Int64']['output']>;
  /** @deprecated Use contractLineItems instead. */
  serviceLineItems?: Maybe<Array<ServiceLineItem>>;
  serviceStarted?: Maybe<Scalars['Time']['output']>;
  /** @deprecated Use serviceStarted instead. */
  serviceStartedAt?: Maybe<Scalars['Time']['output']>;
  /** @deprecated Use contractSigned instead. */
  signedAt?: Maybe<Scalars['Time']['output']>;
  /** @deprecated Use metadata instead. */
  source: DataSource;
  /** @deprecated Use metadata instead. */
  sourceOfTruth: DataSource;
  /** @deprecated Use contractStatus instead. */
  status: ContractStatus;
  /** @deprecated Use metadata instead. */
  updatedAt: Scalars['Time']['output'];
  /** @deprecated Use billingDetails instead. */
  zip?: Maybe<Scalars['String']['output']>;
};

export enum ContractBillingCycle {
  AnnualBilling = 'ANNUAL_BILLING',
  MonthlyBilling = 'MONTHLY_BILLING',
  None = 'NONE',
  QuarterlyBilling = 'QUARTERLY_BILLING',
}

export type ContractInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  billingCycle?: InputMaybe<ContractBillingCycle>;
  billingEnabled?: InputMaybe<Scalars['Boolean']['input']>;
  committedPeriods?: InputMaybe<Scalars['Int64']['input']>;
  contractName?: InputMaybe<Scalars['String']['input']>;
  contractRenewalCycle?: InputMaybe<ContractRenewalCycle>;
  contractSigned?: InputMaybe<Scalars['Time']['input']>;
  contractUrl?: InputMaybe<Scalars['String']['input']>;
  currency?: InputMaybe<Currency>;
  externalReference?: InputMaybe<ExternalSystemReferenceInput>;
  invoicingStartDate?: InputMaybe<Scalars['Time']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  organizationId: Scalars['ID']['input'];
  renewalCycle?: InputMaybe<ContractRenewalCycle>;
  renewalPeriods?: InputMaybe<Scalars['Int64']['input']>;
  serviceStarted?: InputMaybe<Scalars['Time']['input']>;
  serviceStartedAt?: InputMaybe<Scalars['Time']['input']>;
  signedAt?: InputMaybe<Scalars['Time']['input']>;
};

export enum ContractRenewalCycle {
  AnnualRenewal = 'ANNUAL_RENEWAL',
  MonthlyRenewal = 'MONTHLY_RENEWAL',
  None = 'NONE',
  QuarterlyRenewal = 'QUARTERLY_RENEWAL',
}

export enum ContractStatus {
  Draft = 'DRAFT',
  Ended = 'ENDED',
  Live = 'LIVE',
  Undefined = 'UNDEFINED',
}

export type ContractUpdateInput = {
  addressLine1?: InputMaybe<Scalars['String']['input']>;
  addressLine2?: InputMaybe<Scalars['String']['input']>;
  appSource?: InputMaybe<Scalars['String']['input']>;
  billingCycle?: InputMaybe<ContractBillingCycle>;
  billingDetails?: InputMaybe<BillingDetailsInput>;
  billingEnabled?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithBankTransfer?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithCard?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithDirectDebit?: InputMaybe<Scalars['Boolean']['input']>;
  committedPeriods?: InputMaybe<Scalars['Int64']['input']>;
  contractEnded?: InputMaybe<Scalars['Time']['input']>;
  contractId: Scalars['ID']['input'];
  contractName?: InputMaybe<Scalars['String']['input']>;
  contractRenewalCycle?: InputMaybe<ContractRenewalCycle>;
  contractSigned?: InputMaybe<Scalars['Time']['input']>;
  contractUrl?: InputMaybe<Scalars['String']['input']>;
  country?: InputMaybe<Scalars['String']['input']>;
  currency?: InputMaybe<Currency>;
  endedAt?: InputMaybe<Scalars['Time']['input']>;
  invoiceEmail?: InputMaybe<Scalars['String']['input']>;
  invoiceNote?: InputMaybe<Scalars['String']['input']>;
  invoicingStartDate?: InputMaybe<Scalars['Time']['input']>;
  locality?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  organizationLegalName?: InputMaybe<Scalars['String']['input']>;
  patch?: InputMaybe<Scalars['Boolean']['input']>;
  renewalCycle?: InputMaybe<ContractRenewalCycle>;
  renewalPeriods?: InputMaybe<Scalars['Int64']['input']>;
  serviceStarted?: InputMaybe<Scalars['Time']['input']>;
  serviceStartedAt?: InputMaybe<Scalars['Time']['input']>;
  signedAt?: InputMaybe<Scalars['Time']['input']>;
  zip?: InputMaybe<Scalars['String']['input']>;
};

export type Country = {
  __typename?: 'Country';
  codeA2: Scalars['String']['output'];
  codeA3: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  phoneCode: Scalars['String']['output'];
};

export enum Currency {
  Aud = 'AUD',
  Brl = 'BRL',
  Cad = 'CAD',
  Chf = 'CHF',
  Cny = 'CNY',
  Eur = 'EUR',
  Gbp = 'GBP',
  Hkd = 'HKD',
  Inr = 'INR',
  Jpy = 'JPY',
  Krw = 'KRW',
  Mxn = 'MXN',
  Nok = 'NOK',
  Nzd = 'NZD',
  Ron = 'RON',
  Sek = 'SEK',
  Sgd = 'SGD',
  Try = 'TRY',
  Usd = 'USD',
  Zar = 'ZAR',
}

/**
 * Describes a custom, user-defined field associated with a `Contact`.
 * **A `return` object.**
 */
export type CustomField = Node & {
  __typename?: 'CustomField';
  createdAt: Scalars['Time']['output'];
  /**
   * Datatype of the custom field.
   * **Required**
   */
  datatype: CustomFieldDataType;
  /**
   * The unique ID associated with the custom field.
   * **Required**
   */
  id: Scalars['ID']['output'];
  /**
   * The name of the custom field.
   * **Required**
   */
  name: Scalars['String']['output'];
  /** The source of the custom field value */
  source: DataSource;
  template?: Maybe<CustomFieldTemplate>;
  updatedAt: Scalars['Time']['output'];
  /**
   * The value of the custom field.
   * **Required**
   */
  value: Scalars['Any']['output'];
};

export enum CustomFieldDataType {
  Bool = 'BOOL',
  Datetime = 'DATETIME',
  Decimal = 'DECIMAL',
  Integer = 'INTEGER',
  Text = 'TEXT',
}

export type CustomFieldEntityType = {
  entityType: EntityType;
  id: Scalars['ID']['input'];
};

/**
 * Describes a custom, user-defined field associated with a `Contact` of type String.
 * **A `create` object.**
 */
export type CustomFieldInput = {
  /** Datatype of the custom field. */
  datatype?: InputMaybe<CustomFieldDataType>;
  id?: InputMaybe<Scalars['ID']['input']>;
  /** The name of the custom field. */
  name?: InputMaybe<Scalars['String']['input']>;
  templateId?: InputMaybe<Scalars['ID']['input']>;
  /**
   * The value of the custom field.
   * **Required**
   */
  value: Scalars['Any']['input'];
};

export type CustomFieldTemplate = Node & {
  __typename?: 'CustomFieldTemplate';
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  length?: Maybe<Scalars['Int']['output']>;
  mandatory: Scalars['Boolean']['output'];
  max?: Maybe<Scalars['Int']['output']>;
  min?: Maybe<Scalars['Int']['output']>;
  name: Scalars['String']['output'];
  order: Scalars['Int']['output'];
  type: CustomFieldTemplateType;
  updatedAt: Scalars['Time']['output'];
};

export type CustomFieldTemplateInput = {
  length?: InputMaybe<Scalars['Int']['input']>;
  mandatory?: InputMaybe<Scalars['Boolean']['input']>;
  max?: InputMaybe<Scalars['Int']['input']>;
  min?: InputMaybe<Scalars['Int']['input']>;
  name: Scalars['String']['input'];
  order: Scalars['Int']['input'];
  type: CustomFieldTemplateType;
};

export enum CustomFieldTemplateType {
  Link = 'LINK',
  Text = 'TEXT',
}

/**
 * Describes a custom, user-defined field associated with a `Contact`.
 * **An `update` object.**
 */
export type CustomFieldUpdateInput = {
  /**
   * Datatype of the custom field.
   * **Required**
   */
  datatype: CustomFieldDataType;
  /**
   * The unique ID associated with the custom field.
   * **Required**
   */
  id: Scalars['ID']['input'];
  /**
   * The name of the custom field.
   * **Required**
   */
  name: Scalars['String']['input'];
  /**
   * The value of the custom field.
   * **Required**
   */
  value: Scalars['Any']['input'];
};

export type CustomerContact = {
  __typename?: 'CustomerContact';
  email: CustomerEmail;
  id: Scalars['ID']['output'];
};

export type CustomerContactInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  /** An ISO8601 timestamp recording when the contact was created in customerOS. */
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  description?: InputMaybe<Scalars['String']['input']>;
  /** An email addresses associted with the contact. */
  email?: InputMaybe<EmailInput>;
  /** The first name of the contact. */
  firstName?: InputMaybe<Scalars['String']['input']>;
  /** The last name of the contact. */
  lastName?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  /** The prefix of the contact. */
  prefix?: InputMaybe<Scalars['String']['input']>;
  timezone?: InputMaybe<Scalars['String']['input']>;
};

export type CustomerEmail = {
  __typename?: 'CustomerEmail';
  id: Scalars['ID']['output'];
};

export type CustomerJobRole = {
  __typename?: 'CustomerJobRole';
  id: Scalars['ID']['output'];
};

export type CustomerUser = {
  __typename?: 'CustomerUser';
  id: Scalars['ID']['output'];
  jobRole: CustomerJobRole;
};

export type DashboardArrBreakdown = {
  __typename?: 'DashboardARRBreakdown';
  arrBreakdown: Scalars['Float']['output'];
  increasePercentage: Scalars['String']['output'];
  perMonth: Array<Maybe<DashboardArrBreakdownPerMonth>>;
};

export type DashboardArrBreakdownPerMonth = {
  __typename?: 'DashboardARRBreakdownPerMonth';
  cancellations: Scalars['Float']['output'];
  churned: Scalars['Float']['output'];
  downgrades: Scalars['Float']['output'];
  month: Scalars['Int']['output'];
  newlyContracted: Scalars['Float']['output'];
  renewals: Scalars['Float']['output'];
  upsells: Scalars['Float']['output'];
  year: Scalars['Int']['output'];
};

export type DashboardCustomerMap = {
  __typename?: 'DashboardCustomerMap';
  arr: Scalars['Float']['output'];
  contractSignedDate: Scalars['Time']['output'];
  organization: Organization;
  organizationId: Scalars['ID']['output'];
  state: DashboardCustomerMapState;
};

export enum DashboardCustomerMapState {
  AtRisk = 'AT_RISK',
  Churned = 'CHURNED',
  Ok = 'OK',
}

export type DashboardGrossRevenueRetention = {
  __typename?: 'DashboardGrossRevenueRetention';
  grossRevenueRetention: Scalars['Float']['output'];
  /** @deprecated Use increasePercentageValue instead */
  increasePercentage: Scalars['String']['output'];
  increasePercentageValue: Scalars['Float']['output'];
  perMonth: Array<Maybe<DashboardGrossRevenueRetentionPerMonth>>;
};

export type DashboardGrossRevenueRetentionPerMonth = {
  __typename?: 'DashboardGrossRevenueRetentionPerMonth';
  month: Scalars['Int']['output'];
  percentage: Scalars['Float']['output'];
  year: Scalars['Int']['output'];
};

export type DashboardMrrPerCustomer = {
  __typename?: 'DashboardMRRPerCustomer';
  increasePercentage: Scalars['String']['output'];
  mrrPerCustomer: Scalars['Float']['output'];
  perMonth: Array<Maybe<DashboardMrrPerCustomerPerMonth>>;
};

export type DashboardMrrPerCustomerPerMonth = {
  __typename?: 'DashboardMRRPerCustomerPerMonth';
  month: Scalars['Int']['output'];
  value: Scalars['Float']['output'];
  year: Scalars['Int']['output'];
};

export type DashboardNewCustomers = {
  __typename?: 'DashboardNewCustomers';
  perMonth: Array<Maybe<DashboardNewCustomersPerMonth>>;
  thisMonthCount: Scalars['Int']['output'];
  thisMonthIncreasePercentage: Scalars['String']['output'];
};

export type DashboardNewCustomersPerMonth = {
  __typename?: 'DashboardNewCustomersPerMonth';
  count: Scalars['Int']['output'];
  month: Scalars['Int']['output'];
  year: Scalars['Int']['output'];
};

export type DashboardOnboardingCompletion = {
  __typename?: 'DashboardOnboardingCompletion';
  completionPercentage: Scalars['Float']['output'];
  increasePercentage: Scalars['Float']['output'];
  perMonth: Array<DashboardOnboardingCompletionPerMonth>;
};

export type DashboardOnboardingCompletionPerMonth = {
  __typename?: 'DashboardOnboardingCompletionPerMonth';
  month: Scalars['Int']['output'];
  value: Scalars['Float']['output'];
  year: Scalars['Int']['output'];
};

export type DashboardPeriodInput = {
  end: Scalars['Time']['input'];
  start: Scalars['Time']['input'];
};

export type DashboardRetentionRate = {
  __typename?: 'DashboardRetentionRate';
  /** @deprecated Use increasePercentageValue instead */
  increasePercentage: Scalars['String']['output'];
  increasePercentageValue: Scalars['Float']['output'];
  perMonth: Array<Maybe<DashboardRetentionRatePerMonth>>;
  retentionRate: Scalars['Float']['output'];
};

export type DashboardRetentionRatePerMonth = {
  __typename?: 'DashboardRetentionRatePerMonth';
  churnCount: Scalars['Int']['output'];
  month: Scalars['Int']['output'];
  renewCount: Scalars['Int']['output'];
  year: Scalars['Int']['output'];
};

export type DashboardRevenueAtRisk = {
  __typename?: 'DashboardRevenueAtRisk';
  atRisk: Scalars['Float']['output'];
  highConfidence: Scalars['Float']['output'];
};

export type DashboardTimeToOnboard = {
  __typename?: 'DashboardTimeToOnboard';
  increasePercentage?: Maybe<Scalars['Float']['output']>;
  perMonth: Array<DashboardTimeToOnboardPerMonth>;
  timeToOnboard?: Maybe<Scalars['Float']['output']>;
};

export type DashboardTimeToOnboardPerMonth = {
  __typename?: 'DashboardTimeToOnboardPerMonth';
  month: Scalars['Int']['output'];
  value: Scalars['Float']['output'];
  year: Scalars['Int']['output'];
};

export enum DataSource {
  Close = 'CLOSE',
  Hubspot = 'HUBSPOT',
  Intercom = 'INTERCOM',
  Mixpanel = 'MIXPANEL',
  Na = 'NA',
  Openline = 'OPENLINE',
  Outlook = 'OUTLOOK',
  Pipedrive = 'PIPEDRIVE',
  Salesforce = 'SALESFORCE',
  Slack = 'SLACK',
  Stripe = 'STRIPE',
  Unthread = 'UNTHREAD',
  Webscrape = 'WEBSCRAPE',
  ZendeskSupport = 'ZENDESK_SUPPORT',
}

export type DeleteResponse = {
  __typename?: 'DeleteResponse';
  accepted: Scalars['Boolean']['output'];
  completed: Scalars['Boolean']['output'];
};

export type DescriptionNode = InteractionEvent | InteractionSession | Meeting;

/**
 * Describes an email address associated with a `Contact` in customerOS.
 * **A `return` object.**
 */
export type Email = {
  __typename?: 'Email';
  appSource: Scalars['String']['output'];
  contacts: Array<Contact>;
  createdAt: Scalars['Time']['output'];
  /** An email address assocaited with the contact in customerOS. */
  email?: Maybe<Scalars['String']['output']>;
  emailValidationDetails: EmailValidationDetails;
  /**
   * The unique ID associated with the contact in customerOS.
   * **Required**
   */
  id: Scalars['ID']['output'];
  /** Describes the type of email address (WORK, PERSONAL, etc). */
  label?: Maybe<EmailLabel>;
  organizations: Array<Organization>;
  /**
   * Identifies whether the email address is primary or not.
   * **Required.**
   */
  primary: Scalars['Boolean']['output'];
  rawEmail?: Maybe<Scalars['String']['output']>;
  source: DataSource;
  sourceOfTruth: DataSource;
  updatedAt: Scalars['Time']['output'];
  users: Array<User>;
};

/**
 * Describes an email address associated with a `Contact` in customerOS.
 * **A `create` object.**
 */
export type EmailInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  /**
   * An email address associated with the contact in customerOS.
   * **Required.**
   */
  email: Scalars['String']['input'];
  /** Describes the type of email address (WORK, PERSONAL, etc). */
  label?: InputMaybe<EmailLabel>;
  /**
   * Identifies whether the email address is primary or not.
   * **Required.**
   */
  primary?: InputMaybe<Scalars['Boolean']['input']>;
};

/**
 * Describes the type of email address (WORK, PERSONAL, etc).
 * **A `return` object.
 */
export enum EmailLabel {
  Main = 'MAIN',
  Other = 'OTHER',
  Personal = 'PERSONAL',
  Work = 'WORK',
}

export type EmailParticipant = {
  __typename?: 'EmailParticipant';
  emailParticipant: Email;
  type?: Maybe<Scalars['String']['output']>;
};

/**
 * Describes an email address associated with a `Contact` in customerOS.
 * **An `update` object.**
 */
export type EmailUpdateInput = {
  email?: InputMaybe<Scalars['String']['input']>;
  /**
   * An email address assocaited with the contact in customerOS.
   * **Required.**
   */
  id: Scalars['ID']['input'];
  /** Describes the type of email address (WORK, PERSONAL, etc). */
  label?: InputMaybe<EmailLabel>;
  /**
   * Identifies whether the email address is primary or not.
   * **Required.**
   */
  primary?: InputMaybe<Scalars['Boolean']['input']>;
};

export type EmailValidationDetails = {
  __typename?: 'EmailValidationDetails';
  acceptsMail?: Maybe<Scalars['Boolean']['output']>;
  canConnectSmtp?: Maybe<Scalars['Boolean']['output']>;
  error?: Maybe<Scalars['String']['output']>;
  hasFullInbox?: Maybe<Scalars['Boolean']['output']>;
  isCatchAll?: Maybe<Scalars['Boolean']['output']>;
  isDeliverable?: Maybe<Scalars['Boolean']['output']>;
  isDisabled?: Maybe<Scalars['Boolean']['output']>;
  isReachable?: Maybe<Scalars['String']['output']>;
  isValidSyntax?: Maybe<Scalars['Boolean']['output']>;
  validated?: Maybe<Scalars['Boolean']['output']>;
};

export type EntityTemplate = Node & {
  __typename?: 'EntityTemplate';
  createdAt: Scalars['Time']['output'];
  customFieldTemplates: Array<CustomFieldTemplate>;
  extends?: Maybe<EntityTemplateExtension>;
  fieldSetTemplates: Array<FieldSetTemplate>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  updatedAt: Scalars['Time']['output'];
  version: Scalars['Int']['output'];
};

export enum EntityTemplateExtension {
  Contact = 'CONTACT',
  Organization = 'ORGANIZATION',
}

export type EntityTemplateInput = {
  customFieldTemplateInputs?: InputMaybe<Array<CustomFieldTemplateInput>>;
  extends?: InputMaybe<EntityTemplateExtension>;
  fieldSetTemplateInputs?: InputMaybe<Array<FieldSetTemplateInput>>;
  name: Scalars['String']['input'];
};

export enum EntityType {
  Contact = 'Contact',
  Organization = 'Organization',
}

export type ExtensibleEntity = {
  id: Scalars['ID']['output'];
  template?: Maybe<EntityTemplate>;
};

export type ExternalSystem = {
  __typename?: 'ExternalSystem';
  externalId?: Maybe<Scalars['String']['output']>;
  externalSource?: Maybe<Scalars['String']['output']>;
  externalUrl?: Maybe<Scalars['String']['output']>;
  syncDate?: Maybe<Scalars['Time']['output']>;
  type: ExternalSystemType;
};

export type ExternalSystemInstance = {
  __typename?: 'ExternalSystemInstance';
  paymentMethods: Array<Scalars['String']['output']>;
  type: ExternalSystemType;
};

export type ExternalSystemReferenceInput = {
  externalId: Scalars['ID']['input'];
  externalSource?: InputMaybe<Scalars['String']['input']>;
  externalUrl?: InputMaybe<Scalars['String']['input']>;
  syncDate?: InputMaybe<Scalars['Time']['input']>;
  type: ExternalSystemType;
};

export enum ExternalSystemType {
  Calcom = 'CALCOM',
  Close = 'CLOSE',
  Hubspot = 'HUBSPOT',
  Intercom = 'INTERCOM',
  Mixpanel = 'MIXPANEL',
  Outlook = 'OUTLOOK',
  Pipedrive = 'PIPEDRIVE',
  Salesforce = 'SALESFORCE',
  Slack = 'SLACK',
  Stripe = 'STRIPE',
  Unthread = 'UNTHREAD',
  ZendeskSupport = 'ZENDESK_SUPPORT',
}

export type FieldSet = {
  __typename?: 'FieldSet';
  createdAt: Scalars['Time']['output'];
  customFields: Array<CustomField>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  source: DataSource;
  template?: Maybe<FieldSetTemplate>;
  updatedAt: Scalars['Time']['output'];
};

export type FieldSetInput = {
  customFields?: InputMaybe<Array<CustomFieldInput>>;
  id?: InputMaybe<Scalars['ID']['input']>;
  name: Scalars['String']['input'];
  templateId?: InputMaybe<Scalars['ID']['input']>;
};

export type FieldSetTemplate = Node & {
  __typename?: 'FieldSetTemplate';
  createdAt: Scalars['Time']['output'];
  customFieldTemplates: Array<CustomFieldTemplate>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  order: Scalars['Int']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type FieldSetTemplateInput = {
  customFieldTemplateInputs?: InputMaybe<Array<CustomFieldTemplateInput>>;
  name: Scalars['String']['input'];
  order: Scalars['Int']['input'];
};

export type FieldSetUpdateInput = {
  id: Scalars['ID']['input'];
  name: Scalars['String']['input'];
};

export type Filter = {
  AND?: InputMaybe<Array<Filter>>;
  NOT?: InputMaybe<Filter>;
  OR?: InputMaybe<Array<Filter>>;
  filter?: InputMaybe<FilterItem>;
};

export type FilterItem = {
  caseSensitive?: InputMaybe<Scalars['Boolean']['input']>;
  includeEmpty?: InputMaybe<Scalars['Boolean']['input']>;
  operation?: ComparisonOperator;
  property: Scalars['String']['input'];
  value: Scalars['Any']['input'];
};

export enum FundingRound {
  Angel = 'ANGEL',
  Bridge = 'BRIDGE',
  FriendsAndFamily = 'FRIENDS_AND_FAMILY',
  Ipo = 'IPO',
  PreSeed = 'PRE_SEED',
  Seed = 'SEED',
  SeriesA = 'SERIES_A',
  SeriesB = 'SERIES_B',
  SeriesC = 'SERIES_C',
  SeriesD = 'SERIES_D',
  SeriesE = 'SERIES_E',
  SeriesF = 'SERIES_F',
}

export type GCliAttributeKeyValuePair = {
  __typename?: 'GCliAttributeKeyValuePair';
  display?: Maybe<Scalars['String']['output']>;
  key: Scalars['String']['output'];
  value: Scalars['String']['output'];
};

export enum GCliCacheItemType {
  Contact = 'CONTACT',
  Organization = 'ORGANIZATION',
  State = 'STATE',
}

export type GCliItem = {
  __typename?: 'GCliItem';
  data?: Maybe<Array<GCliAttributeKeyValuePair>>;
  display: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  type: GCliSearchResultType;
};

export enum GCliSearchResultType {
  Contact = 'CONTACT',
  Email = 'EMAIL',
  Organization = 'ORGANIZATION',
  OrganizationRelationship = 'ORGANIZATION_RELATIONSHIP',
  State = 'STATE',
}

export type GlobalCache = {
  __typename?: 'GlobalCache';
  cdnLogoUrl: Scalars['String']['output'];
  contractsExist: Scalars['Boolean']['output'];
  gCliCache: Array<GCliItem>;
  isGoogleActive: Scalars['Boolean']['output'];
  isGoogleTokenExpired: Scalars['Boolean']['output'];
  isOwner: Scalars['Boolean']['output'];
  maxARRForecastValue: Scalars['Float']['output'];
  minARRForecastValue: Scalars['Float']['output'];
  user: User;
};

export type InteractionEvent = Node & {
  __typename?: 'InteractionEvent';
  actionItems?: Maybe<Array<ActionItem>>;
  appSource: Scalars['String']['output'];
  channel?: Maybe<Scalars['String']['output']>;
  channelData?: Maybe<Scalars['String']['output']>;
  content?: Maybe<Scalars['String']['output']>;
  contentType?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  eventIdentifier?: Maybe<Scalars['String']['output']>;
  eventType?: Maybe<Scalars['String']['output']>;
  externalLinks: Array<ExternalSystem>;
  id: Scalars['ID']['output'];
  includes: Array<Attachment>;
  interactionSession?: Maybe<InteractionSession>;
  issue?: Maybe<Issue>;
  meeting?: Maybe<Meeting>;
  repliesTo?: Maybe<InteractionEvent>;
  sentBy: Array<InteractionEventParticipant>;
  sentTo: Array<InteractionEventParticipant>;
  source: DataSource;
  sourceOfTruth: DataSource;
  summary?: Maybe<Analysis>;
};

export type InteractionEventInput = {
  appSource: Scalars['String']['input'];
  channel?: InputMaybe<Scalars['String']['input']>;
  channelData?: InputMaybe<Scalars['String']['input']>;
  content?: InputMaybe<Scalars['String']['input']>;
  contentType?: InputMaybe<Scalars['String']['input']>;
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  eventIdentifier?: InputMaybe<Scalars['String']['input']>;
  eventType?: InputMaybe<Scalars['String']['input']>;
  externalId?: InputMaybe<Scalars['String']['input']>;
  externalSystemId?: InputMaybe<Scalars['String']['input']>;
  interactionSession?: InputMaybe<Scalars['ID']['input']>;
  meetingId?: InputMaybe<Scalars['ID']['input']>;
  repliesTo?: InputMaybe<Scalars['ID']['input']>;
  sentBy: Array<InteractionEventParticipantInput>;
  sentTo: Array<InteractionEventParticipantInput>;
};

export type InteractionEventParticipant =
  | ContactParticipant
  | EmailParticipant
  | JobRoleParticipant
  | OrganizationParticipant
  | PhoneNumberParticipant
  | UserParticipant;

export type InteractionEventParticipantInput = {
  contactID?: InputMaybe<Scalars['ID']['input']>;
  email?: InputMaybe<Scalars['String']['input']>;
  phoneNumber?: InputMaybe<Scalars['String']['input']>;
  type?: InputMaybe<Scalars['String']['input']>;
  userID?: InputMaybe<Scalars['ID']['input']>;
};

export type InteractionSession = Node & {
  __typename?: 'InteractionSession';
  appSource: Scalars['String']['output'];
  attendedBy: Array<InteractionSessionParticipant>;
  channel?: Maybe<Scalars['String']['output']>;
  channelData?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  describedBy: Array<Analysis>;
  /** @deprecated Use updatedAt instead */
  endedAt?: Maybe<Scalars['Time']['output']>;
  events: Array<InteractionEvent>;
  id: Scalars['ID']['output'];
  includes: Array<Attachment>;
  name: Scalars['String']['output'];
  sessionIdentifier?: Maybe<Scalars['String']['output']>;
  source: DataSource;
  sourceOfTruth: DataSource;
  /** @deprecated Use createdAt instead */
  startedAt: Scalars['Time']['output'];
  status: Scalars['String']['output'];
  type?: Maybe<Scalars['String']['output']>;
  updatedAt: Scalars['Time']['output'];
};

export type InteractionSessionInput = {
  appSource: Scalars['String']['input'];
  attendedBy?: InputMaybe<Array<InteractionSessionParticipantInput>>;
  channel?: InputMaybe<Scalars['String']['input']>;
  channelData?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  sessionIdentifier?: InputMaybe<Scalars['String']['input']>;
  status: Scalars['String']['input'];
  type?: InputMaybe<Scalars['String']['input']>;
};

export type InteractionSessionParticipant =
  | ContactParticipant
  | EmailParticipant
  | PhoneNumberParticipant
  | UserParticipant;

export type InteractionSessionParticipantInput = {
  contactID?: InputMaybe<Scalars['ID']['input']>;
  email?: InputMaybe<Scalars['String']['input']>;
  phoneNumber?: InputMaybe<Scalars['String']['input']>;
  type?: InputMaybe<Scalars['String']['input']>;
  userID?: InputMaybe<Scalars['ID']['input']>;
};

export enum InternalStage {
  ClosedLost = 'CLOSED_LOST',
  ClosedWon = 'CLOSED_WON',
  Evaluating = 'EVALUATING',
  Open = 'OPEN',
}

export enum InternalType {
  CrossSell = 'CROSS_SELL',
  Nbo = 'NBO',
  Renewal = 'RENEWAL',
  Upsell = 'UPSELL',
}

export type Invoice = MetadataInterface & {
  __typename?: 'Invoice';
  amountDue: Scalars['Float']['output'];
  amountPaid: Scalars['Float']['output'];
  amountRemaining: Scalars['Float']['output'];
  contract: Contract;
  currency: Scalars['String']['output'];
  customer: InvoiceCustomer;
  domesticPaymentsBankInfo?: Maybe<Scalars['String']['output']>;
  dryRun: Scalars['Boolean']['output'];
  due: Scalars['Time']['output'];
  internationalPaymentsBankInfo?: Maybe<Scalars['String']['output']>;
  invoiceLineItems: Array<InvoiceLine>;
  invoiceNumber: Scalars['String']['output'];
  invoicePeriodEnd: Scalars['Time']['output'];
  invoicePeriodStart: Scalars['Time']['output'];
  invoiceUrl: Scalars['String']['output'];
  metadata: Metadata;
  note?: Maybe<Scalars['String']['output']>;
  offCycle: Scalars['Boolean']['output'];
  organization: Organization;
  paid: Scalars['Boolean']['output'];
  postpaid: Scalars['Boolean']['output'];
  provider: InvoiceProvider;
  repositoryFileId: Scalars['String']['output'];
  status?: Maybe<InvoiceStatus>;
  subtotal: Scalars['Float']['output'];
  taxDue: Scalars['Float']['output'];
};

export type InvoiceCustomer = {
  __typename?: 'InvoiceCustomer';
  addressCountry?: Maybe<Scalars['String']['output']>;
  addressLine1?: Maybe<Scalars['String']['output']>;
  addressLine2?: Maybe<Scalars['String']['output']>;
  addressLocality?: Maybe<Scalars['String']['output']>;
  addressZip?: Maybe<Scalars['String']['output']>;
  email?: Maybe<Scalars['String']['output']>;
  name?: Maybe<Scalars['String']['output']>;
};

export type InvoiceLine = MetadataInterface & {
  __typename?: 'InvoiceLine';
  description: Scalars['String']['output'];
  metadata: Metadata;
  price: Scalars['Float']['output'];
  quantity: Scalars['Int']['output'];
  subtotal: Scalars['Float']['output'];
  taxDue: Scalars['Float']['output'];
  total: Scalars['Float']['output'];
};

export type InvoiceLineInput = {
  billed: BilledType;
  name: Scalars['String']['input'];
  price: Scalars['Float']['input'];
  quantity: Scalars['Int']['input'];
  serviceLineItemId?: InputMaybe<Scalars['ID']['input']>;
};

export type InvoiceProvider = {
  __typename?: 'InvoiceProvider';
  addressCountry?: Maybe<Scalars['String']['output']>;
  addressLine1?: Maybe<Scalars['String']['output']>;
  addressLine2?: Maybe<Scalars['String']['output']>;
  addressLocality?: Maybe<Scalars['String']['output']>;
  addressZip?: Maybe<Scalars['String']['output']>;
  logoRepositoryFileId?: Maybe<Scalars['String']['output']>;
  logoUrl?: Maybe<Scalars['String']['output']>;
  name?: Maybe<Scalars['String']['output']>;
};

export type InvoiceSimulateInput = {
  contractId: Scalars['ID']['input'];
  invoiceLines: Array<InvoiceLineInput>;
  periodEndDate?: InputMaybe<Scalars['Time']['input']>;
  periodStartDate?: InputMaybe<Scalars['Time']['input']>;
};

export enum InvoiceStatus {
  Draft = 'DRAFT',
  Due = 'DUE',
  Paid = 'PAID',
  Void = 'VOID',
}

export type InvoiceUpdateInput = {
  id: Scalars['ID']['input'];
  patch: Scalars['Boolean']['input'];
  status?: InputMaybe<InvoiceStatus>;
};

export type InvoicesPage = Pages & {
  __typename?: 'InvoicesPage';
  content: Array<Invoice>;
  totalElements: Scalars['Int64']['output'];
  totalPages: Scalars['Int']['output'];
};

export type InvoicingCycle = Node &
  SourceFields & {
    __typename?: 'InvoicingCycle';
    appSource: Scalars['String']['output'];
    createdAt: Scalars['Time']['output'];
    id: Scalars['ID']['output'];
    source: DataSource;
    sourceOfTruth: DataSource;
    type: InvoicingCycleType;
    updatedAt: Scalars['Time']['output'];
  };

export type InvoicingCycleInput = {
  type: InvoicingCycleType;
};

export enum InvoicingCycleType {
  Anniversary = 'ANNIVERSARY',
  Date = 'DATE',
}

export type InvoicingCycleUpdateInput = {
  id: Scalars['ID']['input'];
  type: InvoicingCycleType;
};

export type Issue = Node &
  SourceFields & {
    __typename?: 'Issue';
    appSource: Scalars['String']['output'];
    assignedTo: Array<IssueParticipant>;
    comments: Array<Comment>;
    createdAt: Scalars['Time']['output'];
    description?: Maybe<Scalars['String']['output']>;
    externalLinks: Array<ExternalSystem>;
    followedBy: Array<IssueParticipant>;
    id: Scalars['ID']['output'];
    interactionEvents: Array<InteractionEvent>;
    priority?: Maybe<Scalars['String']['output']>;
    reportedBy?: Maybe<IssueParticipant>;
    source: DataSource;
    sourceOfTruth: DataSource;
    status: Scalars['String']['output'];
    subject?: Maybe<Scalars['String']['output']>;
    submittedBy?: Maybe<IssueParticipant>;
    tags?: Maybe<Array<Maybe<Tag>>>;
    updatedAt: Scalars['Time']['output'];
  };

export type IssueParticipant =
  | ContactParticipant
  | OrganizationParticipant
  | UserParticipant;

export type IssueSummaryByStatus = {
  __typename?: 'IssueSummaryByStatus';
  count: Scalars['Int64']['output'];
  status: Scalars['String']['output'];
};

/**
 * Describes the relationship a Contact has with a Organization.
 * **A `return` object**
 */
export type JobRole = {
  __typename?: 'JobRole';
  appSource: Scalars['String']['output'];
  company?: Maybe<Scalars['String']['output']>;
  contact?: Maybe<Contact>;
  createdAt: Scalars['Time']['output'];
  description?: Maybe<Scalars['String']['output']>;
  endedAt?: Maybe<Scalars['Time']['output']>;
  id: Scalars['ID']['output'];
  /** The Contact's job title. */
  jobTitle?: Maybe<Scalars['String']['output']>;
  /**
   * Organization associated with a Contact.
   * **Required.**
   */
  organization?: Maybe<Organization>;
  primary: Scalars['Boolean']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
  startedAt?: Maybe<Scalars['Time']['output']>;
  updatedAt: Scalars['Time']['output'];
};

/**
 * Describes the relationship a Contact has with an Organization.
 * **A `create` object**
 */
export type JobRoleInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  company?: InputMaybe<Scalars['String']['input']>;
  description?: InputMaybe<Scalars['String']['input']>;
  endedAt?: InputMaybe<Scalars['Time']['input']>;
  jobTitle?: InputMaybe<Scalars['String']['input']>;
  organizationId?: InputMaybe<Scalars['ID']['input']>;
  primary?: InputMaybe<Scalars['Boolean']['input']>;
  startedAt?: InputMaybe<Scalars['Time']['input']>;
};

export type JobRoleParticipant = {
  __typename?: 'JobRoleParticipant';
  jobRoleParticipant: JobRole;
  type?: Maybe<Scalars['String']['output']>;
};

/**
 * Describes the relationship a Contact has with an Organization.
 * **A `create` object**
 */
export type JobRoleUpdateInput = {
  company?: InputMaybe<Scalars['String']['input']>;
  description?: InputMaybe<Scalars['String']['input']>;
  endedAt?: InputMaybe<Scalars['Time']['input']>;
  id: Scalars['ID']['input'];
  jobTitle?: InputMaybe<Scalars['String']['input']>;
  organizationId?: InputMaybe<Scalars['ID']['input']>;
  primary?: InputMaybe<Scalars['Boolean']['input']>;
  startedAt?: InputMaybe<Scalars['Time']['input']>;
};

export type LastTouchpoint = {
  __typename?: 'LastTouchpoint';
  lastTouchPointAt?: Maybe<Scalars['Time']['output']>;
  lastTouchPointTimelineEvent?: Maybe<TimelineEvent>;
  lastTouchPointTimelineEventId?: Maybe<Scalars['ID']['output']>;
  lastTouchPointType?: Maybe<LastTouchpointType>;
};

export enum LastTouchpointType {
  Action = 'ACTION',
  ActionCreated = 'ACTION_CREATED',
  Analysis = 'ANALYSIS',
  InteractionEventChat = 'INTERACTION_EVENT_CHAT',
  InteractionEventEmailSent = 'INTERACTION_EVENT_EMAIL_SENT',
  InteractionEventPhoneCall = 'INTERACTION_EVENT_PHONE_CALL',
  InteractionSession = 'INTERACTION_SESSION',
  IssueCreated = 'ISSUE_CREATED',
  IssueUpdated = 'ISSUE_UPDATED',
  LogEntry = 'LOG_ENTRY',
  Meeting = 'MEETING',
  Note = 'NOTE',
  PageView = 'PAGE_VIEW',
}

export type LinkOrganizationsInput = {
  organizationId: Scalars['ID']['input'];
  subOrganizationId: Scalars['ID']['input'];
  type?: InputMaybe<Scalars['String']['input']>;
};

export type LinkedOrganization = {
  __typename?: 'LinkedOrganization';
  organization: Organization;
  type?: Maybe<Scalars['String']['output']>;
};

export type Location = Node &
  SourceFields & {
    __typename?: 'Location';
    address?: Maybe<Scalars['String']['output']>;
    address2?: Maybe<Scalars['String']['output']>;
    addressType?: Maybe<Scalars['String']['output']>;
    appSource: Scalars['String']['output'];
    commercial?: Maybe<Scalars['Boolean']['output']>;
    country?: Maybe<Scalars['String']['output']>;
    createdAt: Scalars['Time']['output'];
    district?: Maybe<Scalars['String']['output']>;
    houseNumber?: Maybe<Scalars['String']['output']>;
    id: Scalars['ID']['output'];
    latitude?: Maybe<Scalars['Float']['output']>;
    locality?: Maybe<Scalars['String']['output']>;
    longitude?: Maybe<Scalars['Float']['output']>;
    name?: Maybe<Scalars['String']['output']>;
    plusFour?: Maybe<Scalars['String']['output']>;
    postalCode?: Maybe<Scalars['String']['output']>;
    predirection?: Maybe<Scalars['String']['output']>;
    rawAddress?: Maybe<Scalars['String']['output']>;
    region?: Maybe<Scalars['String']['output']>;
    source: DataSource;
    sourceOfTruth: DataSource;
    street?: Maybe<Scalars['String']['output']>;
    timeZone?: Maybe<Scalars['String']['output']>;
    updatedAt: Scalars['Time']['output'];
    utcOffset?: Maybe<Scalars['Int64']['output']>;
    zip?: Maybe<Scalars['String']['output']>;
  };

export type LocationUpdateInput = {
  address?: InputMaybe<Scalars['String']['input']>;
  address2?: InputMaybe<Scalars['String']['input']>;
  addressType?: InputMaybe<Scalars['String']['input']>;
  commercial?: InputMaybe<Scalars['Boolean']['input']>;
  country?: InputMaybe<Scalars['String']['input']>;
  district?: InputMaybe<Scalars['String']['input']>;
  houseNumber?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['ID']['input'];
  latitude?: InputMaybe<Scalars['Float']['input']>;
  locality?: InputMaybe<Scalars['String']['input']>;
  longitude?: InputMaybe<Scalars['Float']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  plusFour?: InputMaybe<Scalars['String']['input']>;
  postalCode?: InputMaybe<Scalars['String']['input']>;
  predirection?: InputMaybe<Scalars['String']['input']>;
  rawAddress?: InputMaybe<Scalars['String']['input']>;
  region?: InputMaybe<Scalars['String']['input']>;
  street?: InputMaybe<Scalars['String']['input']>;
  timeZone?: InputMaybe<Scalars['String']['input']>;
  utcOffset?: InputMaybe<Scalars['Int64']['input']>;
  zip?: InputMaybe<Scalars['String']['input']>;
};

export type LogEntry = {
  __typename?: 'LogEntry';
  appSource: Scalars['String']['output'];
  content?: Maybe<Scalars['String']['output']>;
  contentType?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  externalLinks: Array<ExternalSystem>;
  id: Scalars['ID']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
  startedAt: Scalars['Time']['output'];
  tags: Array<Tag>;
  updatedAt: Scalars['Time']['output'];
};

export type LogEntryInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  content?: InputMaybe<Scalars['String']['input']>;
  contentType?: InputMaybe<Scalars['String']['input']>;
  startedAt?: InputMaybe<Scalars['Time']['input']>;
  tags?: InputMaybe<Array<TagIdOrNameInput>>;
};

export type LogEntryUpdateInput = {
  content?: InputMaybe<Scalars['String']['input']>;
  contentType?: InputMaybe<Scalars['String']['input']>;
  startedAt?: InputMaybe<Scalars['Time']['input']>;
};

export enum Market {
  B2B = 'B2B',
  B2C = 'B2C',
  Marketplace = 'MARKETPLACE',
}

export type MasterPlan = Node &
  SourceFields & {
    __typename?: 'MasterPlan';
    appSource: Scalars['String']['output'];
    createdAt: Scalars['Time']['output'];
    id: Scalars['ID']['output'];
    milestones: Array<MasterPlanMilestone>;
    name: Scalars['String']['output'];
    retired: Scalars['Boolean']['output'];
    retiredMilestones: Array<MasterPlanMilestone>;
    source: DataSource;
    sourceOfTruth: DataSource;
    updatedAt: Scalars['Time']['output'];
  };

export type MasterPlanInput = {
  name?: InputMaybe<Scalars['String']['input']>;
};

export type MasterPlanMilestone = Node &
  SourceFields & {
    __typename?: 'MasterPlanMilestone';
    appSource: Scalars['String']['output'];
    createdAt: Scalars['Time']['output'];
    durationHours: Scalars['Int64']['output'];
    id: Scalars['ID']['output'];
    items: Array<Scalars['String']['output']>;
    name: Scalars['String']['output'];
    optional: Scalars['Boolean']['output'];
    order: Scalars['Int64']['output'];
    retired: Scalars['Boolean']['output'];
    source: DataSource;
    sourceOfTruth: DataSource;
    updatedAt: Scalars['Time']['output'];
  };

export type MasterPlanMilestoneInput = {
  durationHours: Scalars['Int64']['input'];
  items: Array<Scalars['String']['input']>;
  masterPlanId: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  optional: Scalars['Boolean']['input'];
  order: Scalars['Int64']['input'];
};

export type MasterPlanMilestoneReorderInput = {
  masterPlanId: Scalars['ID']['input'];
  orderedIds: Array<Scalars['ID']['input']>;
};

export type MasterPlanMilestoneUpdateInput = {
  durationHours?: InputMaybe<Scalars['Int64']['input']>;
  id: Scalars['ID']['input'];
  items?: InputMaybe<Array<Scalars['String']['input']>>;
  masterPlanId: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  optional?: InputMaybe<Scalars['Boolean']['input']>;
  order?: InputMaybe<Scalars['Int64']['input']>;
  retired?: InputMaybe<Scalars['Boolean']['input']>;
};

export type MasterPlanUpdateInput = {
  id: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  retired?: InputMaybe<Scalars['Boolean']['input']>;
};

export type Meeting = Node & {
  __typename?: 'Meeting';
  agenda?: Maybe<Scalars['String']['output']>;
  agendaContentType?: Maybe<Scalars['String']['output']>;
  appSource: Scalars['String']['output'];
  attendedBy: Array<MeetingParticipant>;
  conferenceUrl?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy: Array<MeetingParticipant>;
  describedBy: Array<Analysis>;
  endedAt?: Maybe<Scalars['Time']['output']>;
  events: Array<InteractionEvent>;
  externalSystem: Array<ExternalSystem>;
  id: Scalars['ID']['output'];
  includes: Array<Attachment>;
  meetingExternalUrl?: Maybe<Scalars['String']['output']>;
  name?: Maybe<Scalars['String']['output']>;
  note: Array<Note>;
  recording?: Maybe<Attachment>;
  source: DataSource;
  sourceOfTruth: DataSource;
  startedAt?: Maybe<Scalars['Time']['output']>;
  status: MeetingStatus;
  updatedAt: Scalars['Time']['output'];
};

export type MeetingInput = {
  agenda?: InputMaybe<Scalars['String']['input']>;
  agendaContentType?: InputMaybe<Scalars['String']['input']>;
  appSource?: InputMaybe<Scalars['String']['input']>;
  attendedBy?: InputMaybe<Array<MeetingParticipantInput>>;
  conferenceUrl?: InputMaybe<Scalars['String']['input']>;
  createdAt?: InputMaybe<Scalars['Time']['input']>;
  createdBy?: InputMaybe<Array<MeetingParticipantInput>>;
  endedAt?: InputMaybe<Scalars['Time']['input']>;
  externalSystem?: InputMaybe<ExternalSystemReferenceInput>;
  meetingExternalUrl?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  note?: InputMaybe<NoteInput>;
  startedAt?: InputMaybe<Scalars['Time']['input']>;
  status?: InputMaybe<MeetingStatus>;
};

export type MeetingParticipant =
  | ContactParticipant
  | EmailParticipant
  | OrganizationParticipant
  | UserParticipant;

export type MeetingParticipantInput = {
  contactId?: InputMaybe<Scalars['ID']['input']>;
  organizationId?: InputMaybe<Scalars['ID']['input']>;
  userId?: InputMaybe<Scalars['ID']['input']>;
};

export enum MeetingStatus {
  Accepted = 'ACCEPTED',
  Canceled = 'CANCELED',
  Undefined = 'UNDEFINED',
}

export type MeetingUpdateInput = {
  agenda?: InputMaybe<Scalars['String']['input']>;
  agendaContentType?: InputMaybe<Scalars['String']['input']>;
  appSource?: InputMaybe<Scalars['String']['input']>;
  conferenceUrl?: InputMaybe<Scalars['String']['input']>;
  endedAt?: InputMaybe<Scalars['Time']['input']>;
  externalSystem?: InputMaybe<ExternalSystemReferenceInput>;
  meetingExternalUrl?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  note?: InputMaybe<NoteUpdateInput>;
  startedAt?: InputMaybe<Scalars['Time']['input']>;
  status?: InputMaybe<MeetingStatus>;
};

/**
 * Specifies how many pages of meeting information has been returned in the query response.
 * **A `response` object.**
 */
export type MeetingsPage = Pages & {
  __typename?: 'MeetingsPage';
  /**
   * A contact entity in customerOS.
   * **Required.  If no values it returns an empty array.**
   */
  content: Array<Meeting>;
  /**
   * Total number of elements in the query response.
   * **Required.**
   */
  totalElements: Scalars['Int64']['output'];
  /**
   * Total number of pages in the query response.
   * **Required.**
   */
  totalPages: Scalars['Int']['output'];
};

export type Metadata = Node &
  SourceFieldsInterface & {
    __typename?: 'Metadata';
    appSource: Scalars['String']['output'];
    created: Scalars['Time']['output'];
    id: Scalars['ID']['output'];
    lastUpdated: Scalars['Time']['output'];
    source: DataSource;
    sourceOfTruth: DataSource;
  };

export type MetadataInterface = {
  metadata: Metadata;
};

export type Mutation = {
  __typename?: 'Mutation';
  analysis_Create: Analysis;
  attachment_Create: Attachment;
  bankAccount_Create: BankAccount;
  bankAccount_Delete: DeleteResponse;
  bankAccount_Update: BankAccount;
  billingProfile_Create: Scalars['ID']['output'];
  billingProfile_LinkEmail: Scalars['ID']['output'];
  billingProfile_LinkLocation: Scalars['ID']['output'];
  billingProfile_UnlinkEmail: Scalars['ID']['output'];
  billingProfile_UnlinkLocation: Scalars['ID']['output'];
  billingProfile_Update: Scalars['ID']['output'];
  contact_AddNewLocation: Location;
  contact_AddOrganizationById: Contact;
  contact_AddSocial: Social;
  contact_AddTagById: Contact;
  contact_Archive: Result;
  contact_Create: Contact;
  contact_HardDelete: Result;
  contact_Merge: Contact;
  contact_RemoveLocation: Contact;
  contact_RemoveOrganizationById: Contact;
  contact_RemoveTagById: Contact;
  contact_RestoreFromArchive: Result;
  contact_Update: Contact;
  contractLineItem_Close: Scalars['ID']['output'];
  contractLineItem_Create: ServiceLineItem;
  contractLineItem_Update: ServiceLineItem;
  contract_Create: Contract;
  contract_Delete: DeleteResponse;
  contract_Update: Contract;
  customFieldDeleteFromContactById: Result;
  customFieldDeleteFromContactByName: Result;
  customFieldDeleteFromFieldSetById: Result;
  customFieldMergeToContact: CustomField;
  customFieldMergeToFieldSet: CustomField;
  customFieldTemplate_Create: CustomFieldTemplate;
  customFieldUpdateInContact: CustomField;
  customFieldUpdateInFieldSet: CustomField;
  customFieldsMergeAndUpdateInContact: Contact;
  customer_contact_Create: CustomerContact;
  customer_user_AddJobRole: CustomerUser;
  emailDelete: Result;
  emailMergeToContact: Email;
  emailMergeToOrganization: Email;
  emailMergeToUser: Email;
  emailRemoveFromContact: Result;
  emailRemoveFromContactById: Result;
  emailRemoveFromOrganization: Result;
  emailRemoveFromOrganizationById: Result;
  emailRemoveFromUser: Result;
  emailRemoveFromUserById: Result;
  emailUpdateInContact: Email;
  emailUpdateInOrganization: Email;
  emailUpdateInUser: Email;
  entityTemplateCreate: EntityTemplate;
  fieldSetDeleteFromContact: Result;
  fieldSetMergeToContact?: Maybe<FieldSet>;
  fieldSetUpdateInContact?: Maybe<FieldSet>;
  interactionEvent_Create: InteractionEvent;
  interactionEvent_LinkAttachment: InteractionEvent;
  interactionSession_Create: InteractionSession;
  interactionSession_LinkAttachment: InteractionSession;
  invoice_NextDryRunForContract: Scalars['ID']['output'];
  invoice_Pay: Invoice;
  invoice_Simulate: Scalars['ID']['output'];
  invoice_Update: Invoice;
  invoice_Void: Invoice;
  invoicingCycle_Create: InvoicingCycle;
  invoicingCycle_Update: InvoicingCycle;
  jobRole_Create: JobRole;
  jobRole_Delete: Result;
  jobRole_Update: JobRole;
  location_RemoveFromContact: Contact;
  location_RemoveFromOrganization: Organization;
  location_Update: Location;
  logEntry_AddTag: Scalars['ID']['output'];
  logEntry_CreateForOrganization: Scalars['ID']['output'];
  logEntry_RemoveTag: Scalars['ID']['output'];
  logEntry_ResetTags: Scalars['ID']['output'];
  logEntry_Update: Scalars['ID']['output'];
  masterPlanMilestone_BulkUpdate: Array<MasterPlanMilestone>;
  masterPlanMilestone_Create: MasterPlanMilestone;
  masterPlanMilestone_Duplicate: MasterPlanMilestone;
  masterPlanMilestone_Reorder: Scalars['ID']['output'];
  masterPlanMilestone_Update: MasterPlanMilestone;
  masterPlan_Create: MasterPlan;
  masterPlan_CreateDefault: MasterPlan;
  masterPlan_Duplicate: MasterPlan;
  masterPlan_Update: MasterPlan;
  meeting_AddNewLocation: Location;
  meeting_AddNote: Meeting;
  meeting_Create: Meeting;
  meeting_LinkAttachment: Meeting;
  meeting_LinkAttendedBy: Meeting;
  meeting_LinkRecording: Meeting;
  meeting_UnlinkAttachment: Meeting;
  meeting_UnlinkAttendedBy: Meeting;
  meeting_UnlinkRecording: Meeting;
  meeting_Update: Meeting;
  note_CreateForContact: Note;
  note_CreateForOrganization: Note;
  note_Delete: Result;
  note_LinkAttachment: Note;
  note_UnlinkAttachment: Note;
  note_Update: Note;
  opportunityRenewalUpdate: Opportunity;
  opportunityUpdate: Opportunity;
  organizationPlanMilestone_BulkUpdate: Array<OrganizationPlanMilestone>;
  organizationPlanMilestone_Create: OrganizationPlanMilestone;
  organizationPlanMilestone_Duplicate: OrganizationPlanMilestone;
  organizationPlanMilestone_Reorder: Scalars['ID']['output'];
  organizationPlanMilestone_Update: OrganizationPlanMilestone;
  organizationPlan_Create: OrganizationPlan;
  organizationPlan_Duplicate: OrganizationPlan;
  organizationPlan_Update: OrganizationPlan;
  organization_AddNewLocation: Location;
  organization_AddSocial: Social;
  organization_AddSubsidiary: Organization;
  organization_Archive?: Maybe<Result>;
  organization_ArchiveAll?: Maybe<Result>;
  organization_Create: Organization;
  organization_Hide: Scalars['ID']['output'];
  organization_HideAll?: Maybe<Result>;
  organization_Merge: Organization;
  organization_RemoveSubsidiary: Organization;
  organization_SetOwner: Organization;
  organization_Show: Scalars['ID']['output'];
  organization_ShowAll?: Maybe<Result>;
  organization_UnsetOwner: Organization;
  organization_Update: Organization;
  organization_UpdateOnboardingStatus: Organization;
  phoneNumberMergeToContact: PhoneNumber;
  phoneNumberMergeToOrganization: PhoneNumber;
  phoneNumberMergeToUser: PhoneNumber;
  phoneNumberRemoveFromContactByE164: Result;
  phoneNumberRemoveFromContactById: Result;
  phoneNumberRemoveFromOrganizationByE164: Result;
  phoneNumberRemoveFromOrganizationById: Result;
  phoneNumberRemoveFromUserByE164: Result;
  phoneNumberRemoveFromUserById: Result;
  phoneNumberUpdateInContact: PhoneNumber;
  phoneNumberUpdateInOrganization: PhoneNumber;
  phoneNumberUpdateInUser: PhoneNumber;
  player_Merge: Result;
  reminder_Create: Reminder;
  reminder_Update: Reminder;
  serviceLineItem_BulkUpdate: Array<Scalars['ID']['output']>;
  serviceLineItem_Delete: DeleteResponse;
  social_Remove: Result;
  social_Update: Social;
  tag_Create: Tag;
  tag_Delete?: Maybe<Result>;
  tag_Update?: Maybe<Tag>;
  tenant_AddBillingProfile: TenantBillingProfile;
  tenant_Merge: Scalars['String']['output'];
  tenant_UpdateBillingProfile: TenantBillingProfile;
  tenant_UpdateSettings: TenantSettings;
  user_AddRole: User;
  user_AddRoleInTenant: User;
  user_Create: User;
  user_Delete: Result;
  user_DeleteInTenant: Result;
  user_RemoveRole: User;
  user_RemoveRoleInTenant: User;
  user_Update: User;
  workspace_Merge: Result;
  workspace_MergeToTenant: Result;
};

export type MutationAnalysis_CreateArgs = {
  analysis: AnalysisInput;
};

export type MutationAttachment_CreateArgs = {
  input: AttachmentInput;
};

export type MutationBankAccount_CreateArgs = {
  input?: InputMaybe<BankAccountCreateInput>;
};

export type MutationBankAccount_DeleteArgs = {
  id: Scalars['ID']['input'];
};

export type MutationBankAccount_UpdateArgs = {
  input?: InputMaybe<BankAccountUpdateInput>;
};

export type MutationBillingProfile_CreateArgs = {
  input: BillingProfileInput;
};

export type MutationBillingProfile_LinkEmailArgs = {
  input: BillingProfileLinkEmailInput;
};

export type MutationBillingProfile_LinkLocationArgs = {
  input: BillingProfileLinkLocationInput;
};

export type MutationBillingProfile_UnlinkEmailArgs = {
  input: BillingProfileLinkEmailInput;
};

export type MutationBillingProfile_UnlinkLocationArgs = {
  input: BillingProfileLinkLocationInput;
};

export type MutationBillingProfile_UpdateArgs = {
  input: BillingProfileUpdateInput;
};

export type MutationContact_AddNewLocationArgs = {
  contactId: Scalars['ID']['input'];
};

export type MutationContact_AddOrganizationByIdArgs = {
  input: ContactOrganizationInput;
};

export type MutationContact_AddSocialArgs = {
  contactId: Scalars['ID']['input'];
  input: SocialInput;
};

export type MutationContact_AddTagByIdArgs = {
  input: ContactTagInput;
};

export type MutationContact_ArchiveArgs = {
  contactId: Scalars['ID']['input'];
};

export type MutationContact_CreateArgs = {
  input: ContactInput;
};

export type MutationContact_HardDeleteArgs = {
  contactId: Scalars['ID']['input'];
};

export type MutationContact_MergeArgs = {
  mergedContactIds: Array<Scalars['ID']['input']>;
  primaryContactId: Scalars['ID']['input'];
};

export type MutationContact_RemoveLocationArgs = {
  contactId: Scalars['ID']['input'];
  locationId: Scalars['ID']['input'];
};

export type MutationContact_RemoveOrganizationByIdArgs = {
  input: ContactOrganizationInput;
};

export type MutationContact_RemoveTagByIdArgs = {
  input: ContactTagInput;
};

export type MutationContact_RestoreFromArchiveArgs = {
  contactId: Scalars['ID']['input'];
};

export type MutationContact_UpdateArgs = {
  input: ContactUpdateInput;
};

export type MutationContractLineItem_CloseArgs = {
  input: ServiceLineItemCloseInput;
};

export type MutationContractLineItem_CreateArgs = {
  input: ServiceLineItemInput;
};

export type MutationContractLineItem_UpdateArgs = {
  input: ServiceLineItemUpdateInput;
};

export type MutationContract_CreateArgs = {
  input: ContractInput;
};

export type MutationContract_DeleteArgs = {
  id: Scalars['ID']['input'];
};

export type MutationContract_UpdateArgs = {
  input: ContractUpdateInput;
};

export type MutationCustomFieldDeleteFromContactByIdArgs = {
  contactId: Scalars['ID']['input'];
  id: Scalars['ID']['input'];
};

export type MutationCustomFieldDeleteFromContactByNameArgs = {
  contactId: Scalars['ID']['input'];
  fieldName: Scalars['String']['input'];
};

export type MutationCustomFieldDeleteFromFieldSetByIdArgs = {
  contactId: Scalars['ID']['input'];
  fieldSetId: Scalars['ID']['input'];
  id: Scalars['ID']['input'];
};

export type MutationCustomFieldMergeToContactArgs = {
  contactId: Scalars['ID']['input'];
  input: CustomFieldInput;
};

export type MutationCustomFieldMergeToFieldSetArgs = {
  contactId: Scalars['ID']['input'];
  fieldSetId: Scalars['ID']['input'];
  input: CustomFieldInput;
};

export type MutationCustomFieldTemplate_CreateArgs = {
  input: CustomFieldTemplateInput;
};

export type MutationCustomFieldUpdateInContactArgs = {
  contactId: Scalars['ID']['input'];
  input: CustomFieldUpdateInput;
};

export type MutationCustomFieldUpdateInFieldSetArgs = {
  contactId: Scalars['ID']['input'];
  fieldSetId: Scalars['ID']['input'];
  input: CustomFieldUpdateInput;
};

export type MutationCustomFieldsMergeAndUpdateInContactArgs = {
  contactId: Scalars['ID']['input'];
  customFields?: InputMaybe<Array<CustomFieldInput>>;
  fieldSets?: InputMaybe<Array<FieldSetInput>>;
};

export type MutationCustomer_Contact_CreateArgs = {
  input: CustomerContactInput;
};

export type MutationCustomer_User_AddJobRoleArgs = {
  id: Scalars['ID']['input'];
  jobRoleInput: JobRoleInput;
};

export type MutationEmailDeleteArgs = {
  id: Scalars['ID']['input'];
};

export type MutationEmailMergeToContactArgs = {
  contactId: Scalars['ID']['input'];
  input: EmailInput;
};

export type MutationEmailMergeToOrganizationArgs = {
  input: EmailInput;
  organizationId: Scalars['ID']['input'];
};

export type MutationEmailMergeToUserArgs = {
  input: EmailInput;
  userId: Scalars['ID']['input'];
};

export type MutationEmailRemoveFromContactArgs = {
  contactId: Scalars['ID']['input'];
  email: Scalars['String']['input'];
};

export type MutationEmailRemoveFromContactByIdArgs = {
  contactId: Scalars['ID']['input'];
  id: Scalars['ID']['input'];
};

export type MutationEmailRemoveFromOrganizationArgs = {
  email: Scalars['String']['input'];
  organizationId: Scalars['ID']['input'];
};

export type MutationEmailRemoveFromOrganizationByIdArgs = {
  id: Scalars['ID']['input'];
  organizationId: Scalars['ID']['input'];
};

export type MutationEmailRemoveFromUserArgs = {
  email: Scalars['String']['input'];
  userId: Scalars['ID']['input'];
};

export type MutationEmailRemoveFromUserByIdArgs = {
  id: Scalars['ID']['input'];
  userId: Scalars['ID']['input'];
};

export type MutationEmailUpdateInContactArgs = {
  contactId: Scalars['ID']['input'];
  input: EmailUpdateInput;
};

export type MutationEmailUpdateInOrganizationArgs = {
  input: EmailUpdateInput;
  organizationId: Scalars['ID']['input'];
};

export type MutationEmailUpdateInUserArgs = {
  input: EmailUpdateInput;
  userId: Scalars['ID']['input'];
};

export type MutationEntityTemplateCreateArgs = {
  input: EntityTemplateInput;
};

export type MutationFieldSetDeleteFromContactArgs = {
  contactId: Scalars['ID']['input'];
  id: Scalars['ID']['input'];
};

export type MutationFieldSetMergeToContactArgs = {
  contactId: Scalars['ID']['input'];
  input: FieldSetInput;
};

export type MutationFieldSetUpdateInContactArgs = {
  contactId: Scalars['ID']['input'];
  input: FieldSetUpdateInput;
};

export type MutationInteractionEvent_CreateArgs = {
  event: InteractionEventInput;
};

export type MutationInteractionEvent_LinkAttachmentArgs = {
  attachmentId: Scalars['ID']['input'];
  eventId: Scalars['ID']['input'];
};

export type MutationInteractionSession_CreateArgs = {
  session: InteractionSessionInput;
};

export type MutationInteractionSession_LinkAttachmentArgs = {
  attachmentId: Scalars['ID']['input'];
  sessionId: Scalars['ID']['input'];
};

export type MutationInvoice_NextDryRunForContractArgs = {
  contractId: Scalars['ID']['input'];
};

export type MutationInvoice_PayArgs = {
  id: Scalars['ID']['input'];
};

export type MutationInvoice_SimulateArgs = {
  input: InvoiceSimulateInput;
};

export type MutationInvoice_UpdateArgs = {
  input: InvoiceUpdateInput;
};

export type MutationInvoice_VoidArgs = {
  id: Scalars['ID']['input'];
};

export type MutationInvoicingCycle_CreateArgs = {
  input: InvoicingCycleInput;
};

export type MutationInvoicingCycle_UpdateArgs = {
  input: InvoicingCycleUpdateInput;
};

export type MutationJobRole_CreateArgs = {
  contactId: Scalars['ID']['input'];
  input: JobRoleInput;
};

export type MutationJobRole_DeleteArgs = {
  contactId: Scalars['ID']['input'];
  roleId: Scalars['ID']['input'];
};

export type MutationJobRole_UpdateArgs = {
  contactId: Scalars['ID']['input'];
  input: JobRoleUpdateInput;
};

export type MutationLocation_RemoveFromContactArgs = {
  contactId: Scalars['ID']['input'];
  locationId: Scalars['ID']['input'];
};

export type MutationLocation_RemoveFromOrganizationArgs = {
  locationId: Scalars['ID']['input'];
  organizationId: Scalars['ID']['input'];
};

export type MutationLocation_UpdateArgs = {
  input: LocationUpdateInput;
};

export type MutationLogEntry_AddTagArgs = {
  id: Scalars['ID']['input'];
  input: TagIdOrNameInput;
};

export type MutationLogEntry_CreateForOrganizationArgs = {
  input: LogEntryInput;
  organizationId: Scalars['ID']['input'];
};

export type MutationLogEntry_RemoveTagArgs = {
  id: Scalars['ID']['input'];
  input: TagIdOrNameInput;
};

export type MutationLogEntry_ResetTagsArgs = {
  id: Scalars['ID']['input'];
  input?: InputMaybe<Array<TagIdOrNameInput>>;
};

export type MutationLogEntry_UpdateArgs = {
  id: Scalars['ID']['input'];
  input: LogEntryUpdateInput;
};

export type MutationMasterPlanMilestone_BulkUpdateArgs = {
  input: Array<MasterPlanMilestoneUpdateInput>;
};

export type MutationMasterPlanMilestone_CreateArgs = {
  input: MasterPlanMilestoneInput;
};

export type MutationMasterPlanMilestone_DuplicateArgs = {
  id: Scalars['ID']['input'];
  masterPlanId: Scalars['ID']['input'];
};

export type MutationMasterPlanMilestone_ReorderArgs = {
  input: MasterPlanMilestoneReorderInput;
};

export type MutationMasterPlanMilestone_UpdateArgs = {
  input: MasterPlanMilestoneUpdateInput;
};

export type MutationMasterPlan_CreateArgs = {
  input: MasterPlanInput;
};

export type MutationMasterPlan_CreateDefaultArgs = {
  input: MasterPlanInput;
};

export type MutationMasterPlan_DuplicateArgs = {
  id: Scalars['ID']['input'];
};

export type MutationMasterPlan_UpdateArgs = {
  input: MasterPlanUpdateInput;
};

export type MutationMeeting_AddNewLocationArgs = {
  meetingId: Scalars['ID']['input'];
};

export type MutationMeeting_AddNoteArgs = {
  meetingId: Scalars['ID']['input'];
  note?: InputMaybe<NoteInput>;
};

export type MutationMeeting_CreateArgs = {
  meeting: MeetingInput;
};

export type MutationMeeting_LinkAttachmentArgs = {
  attachmentId: Scalars['ID']['input'];
  meetingId: Scalars['ID']['input'];
};

export type MutationMeeting_LinkAttendedByArgs = {
  meetingId: Scalars['ID']['input'];
  participant: MeetingParticipantInput;
};

export type MutationMeeting_LinkRecordingArgs = {
  attachmentId: Scalars['ID']['input'];
  meetingId: Scalars['ID']['input'];
};

export type MutationMeeting_UnlinkAttachmentArgs = {
  attachmentId: Scalars['ID']['input'];
  meetingId: Scalars['ID']['input'];
};

export type MutationMeeting_UnlinkAttendedByArgs = {
  meetingId: Scalars['ID']['input'];
  participant: MeetingParticipantInput;
};

export type MutationMeeting_UnlinkRecordingArgs = {
  attachmentId: Scalars['ID']['input'];
  meetingId: Scalars['ID']['input'];
};

export type MutationMeeting_UpdateArgs = {
  meeting: MeetingUpdateInput;
  meetingId: Scalars['ID']['input'];
};

export type MutationNote_CreateForContactArgs = {
  contactId: Scalars['ID']['input'];
  input: NoteInput;
};

export type MutationNote_CreateForOrganizationArgs = {
  input: NoteInput;
  organizationId: Scalars['ID']['input'];
};

export type MutationNote_DeleteArgs = {
  id: Scalars['ID']['input'];
};

export type MutationNote_LinkAttachmentArgs = {
  attachmentId: Scalars['ID']['input'];
  noteId: Scalars['ID']['input'];
};

export type MutationNote_UnlinkAttachmentArgs = {
  attachmentId: Scalars['ID']['input'];
  noteId: Scalars['ID']['input'];
};

export type MutationNote_UpdateArgs = {
  input: NoteUpdateInput;
};

export type MutationOpportunityRenewalUpdateArgs = {
  input: OpportunityRenewalUpdateInput;
  ownerUserId?: InputMaybe<Scalars['ID']['input']>;
};

export type MutationOpportunityUpdateArgs = {
  input: OpportunityUpdateInput;
};

export type MutationOrganizationPlanMilestone_BulkUpdateArgs = {
  input: Array<OrganizationPlanMilestoneUpdateInput>;
};

export type MutationOrganizationPlanMilestone_CreateArgs = {
  input: OrganizationPlanMilestoneInput;
};

export type MutationOrganizationPlanMilestone_DuplicateArgs = {
  id: Scalars['ID']['input'];
  organizationId: Scalars['ID']['input'];
  organizationPlanId: Scalars['ID']['input'];
};

export type MutationOrganizationPlanMilestone_ReorderArgs = {
  input: OrganizationPlanMilestoneReorderInput;
};

export type MutationOrganizationPlanMilestone_UpdateArgs = {
  input: OrganizationPlanMilestoneUpdateInput;
};

export type MutationOrganizationPlan_CreateArgs = {
  input: OrganizationPlanInput;
};

export type MutationOrganizationPlan_DuplicateArgs = {
  id: Scalars['ID']['input'];
  organizationId: Scalars['ID']['input'];
};

export type MutationOrganizationPlan_UpdateArgs = {
  input: OrganizationPlanUpdateInput;
};

export type MutationOrganization_AddNewLocationArgs = {
  organizationId: Scalars['ID']['input'];
};

export type MutationOrganization_AddSocialArgs = {
  input: SocialInput;
  organizationId: Scalars['ID']['input'];
};

export type MutationOrganization_AddSubsidiaryArgs = {
  input: LinkOrganizationsInput;
};

export type MutationOrganization_ArchiveArgs = {
  id: Scalars['ID']['input'];
};

export type MutationOrganization_ArchiveAllArgs = {
  ids: Array<Scalars['ID']['input']>;
};

export type MutationOrganization_CreateArgs = {
  input: OrganizationInput;
};

export type MutationOrganization_HideArgs = {
  id: Scalars['ID']['input'];
};

export type MutationOrganization_HideAllArgs = {
  ids: Array<Scalars['ID']['input']>;
};

export type MutationOrganization_MergeArgs = {
  mergedOrganizationIds: Array<Scalars['ID']['input']>;
  primaryOrganizationId: Scalars['ID']['input'];
};

export type MutationOrganization_RemoveSubsidiaryArgs = {
  organizationId: Scalars['ID']['input'];
  subsidiaryId: Scalars['ID']['input'];
};

export type MutationOrganization_SetOwnerArgs = {
  organizationId: Scalars['ID']['input'];
  userId: Scalars['ID']['input'];
};

export type MutationOrganization_ShowArgs = {
  id: Scalars['ID']['input'];
};

export type MutationOrganization_ShowAllArgs = {
  ids: Array<Scalars['ID']['input']>;
};

export type MutationOrganization_UnsetOwnerArgs = {
  organizationId: Scalars['ID']['input'];
};

export type MutationOrganization_UpdateArgs = {
  input: OrganizationUpdateInput;
};

export type MutationOrganization_UpdateOnboardingStatusArgs = {
  input: OnboardingStatusInput;
};

export type MutationPhoneNumberMergeToContactArgs = {
  contactId: Scalars['ID']['input'];
  input: PhoneNumberInput;
};

export type MutationPhoneNumberMergeToOrganizationArgs = {
  input: PhoneNumberInput;
  organizationId: Scalars['ID']['input'];
};

export type MutationPhoneNumberMergeToUserArgs = {
  input: PhoneNumberInput;
  userId: Scalars['ID']['input'];
};

export type MutationPhoneNumberRemoveFromContactByE164Args = {
  contactId: Scalars['ID']['input'];
  e164: Scalars['String']['input'];
};

export type MutationPhoneNumberRemoveFromContactByIdArgs = {
  contactId: Scalars['ID']['input'];
  id: Scalars['ID']['input'];
};

export type MutationPhoneNumberRemoveFromOrganizationByE164Args = {
  e164: Scalars['String']['input'];
  organizationId: Scalars['ID']['input'];
};

export type MutationPhoneNumberRemoveFromOrganizationByIdArgs = {
  id: Scalars['ID']['input'];
  organizationId: Scalars['ID']['input'];
};

export type MutationPhoneNumberRemoveFromUserByE164Args = {
  e164: Scalars['String']['input'];
  userId: Scalars['ID']['input'];
};

export type MutationPhoneNumberRemoveFromUserByIdArgs = {
  id: Scalars['ID']['input'];
  userId: Scalars['ID']['input'];
};

export type MutationPhoneNumberUpdateInContactArgs = {
  contactId: Scalars['ID']['input'];
  input: PhoneNumberUpdateInput;
};

export type MutationPhoneNumberUpdateInOrganizationArgs = {
  input: PhoneNumberUpdateInput;
  organizationId: Scalars['ID']['input'];
};

export type MutationPhoneNumberUpdateInUserArgs = {
  input: PhoneNumberUpdateInput;
  userId: Scalars['ID']['input'];
};

export type MutationPlayer_MergeArgs = {
  input: PlayerInput;
  userId: Scalars['ID']['input'];
};

export type MutationReminder_CreateArgs = {
  input: ReminderInput;
};

export type MutationReminder_UpdateArgs = {
  input: ReminderUpdateInput;
};

export type MutationServiceLineItem_BulkUpdateArgs = {
  input: ServiceLineItemBulkUpdateInput;
};

export type MutationServiceLineItem_DeleteArgs = {
  id: Scalars['ID']['input'];
};

export type MutationSocial_RemoveArgs = {
  socialId: Scalars['ID']['input'];
};

export type MutationSocial_UpdateArgs = {
  input: SocialUpdateInput;
};

export type MutationTag_CreateArgs = {
  input: TagInput;
};

export type MutationTag_DeleteArgs = {
  id: Scalars['ID']['input'];
};

export type MutationTag_UpdateArgs = {
  input: TagUpdateInput;
};

export type MutationTenant_AddBillingProfileArgs = {
  input: TenantBillingProfileInput;
};

export type MutationTenant_MergeArgs = {
  tenant: TenantInput;
};

export type MutationTenant_UpdateBillingProfileArgs = {
  input: TenantBillingProfileUpdateInput;
};

export type MutationTenant_UpdateSettingsArgs = {
  input?: InputMaybe<TenantSettingsInput>;
};

export type MutationUser_AddRoleArgs = {
  id: Scalars['ID']['input'];
  role: Role;
};

export type MutationUser_AddRoleInTenantArgs = {
  id: Scalars['ID']['input'];
  role: Role;
  tenant: Scalars['String']['input'];
};

export type MutationUser_CreateArgs = {
  input: UserInput;
};

export type MutationUser_DeleteArgs = {
  id: Scalars['ID']['input'];
};

export type MutationUser_DeleteInTenantArgs = {
  id: Scalars['ID']['input'];
  tenant: Scalars['String']['input'];
};

export type MutationUser_RemoveRoleArgs = {
  id: Scalars['ID']['input'];
  role: Role;
};

export type MutationUser_RemoveRoleInTenantArgs = {
  id: Scalars['ID']['input'];
  role: Role;
  tenant: Scalars['String']['input'];
};

export type MutationUser_UpdateArgs = {
  input: UserUpdateInput;
};

export type MutationWorkspace_MergeArgs = {
  workspace: WorkspaceInput;
};

export type MutationWorkspace_MergeToTenantArgs = {
  tenant: Scalars['String']['input'];
  workspace: WorkspaceInput;
};

export type Node = {
  id: Scalars['ID']['output'];
};

export type Note = {
  __typename?: 'Note';
  appSource: Scalars['String']['output'];
  content?: Maybe<Scalars['String']['output']>;
  contentType?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  id: Scalars['ID']['output'];
  includes: Array<Attachment>;
  noted: Array<NotedEntity>;
  source: DataSource;
  sourceOfTruth: DataSource;
  updatedAt: Scalars['Time']['output'];
};

export type NoteInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  content?: InputMaybe<Scalars['String']['input']>;
  contentType?: InputMaybe<Scalars['String']['input']>;
};

export type NotePage = Pages & {
  __typename?: 'NotePage';
  content: Array<Note>;
  totalElements: Scalars['Int64']['output'];
  totalPages: Scalars['Int']['output'];
};

export type NoteUpdateInput = {
  content?: InputMaybe<Scalars['String']['input']>;
  contentType?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['ID']['input'];
};

export type NotedEntity = Contact | Organization;

export type OnboardingDetails = {
  __typename?: 'OnboardingDetails';
  comments?: Maybe<Scalars['String']['output']>;
  status: OnboardingStatus;
  updatedAt?: Maybe<Scalars['Time']['output']>;
};

export enum OnboardingPlanMilestoneItemStatus {
  Done = 'DONE',
  DoneLate = 'DONE_LATE',
  NotDone = 'NOT_DONE',
  NotDoneLate = 'NOT_DONE_LATE',
  Skipped = 'SKIPPED',
  SkippedLate = 'SKIPPED_LATE',
}

export enum OnboardingPlanMilestoneStatus {
  Done = 'DONE',
  DoneLate = 'DONE_LATE',
  NotStarted = 'NOT_STARTED',
  NotStartedLate = 'NOT_STARTED_LATE',
  Started = 'STARTED',
  StartedLate = 'STARTED_LATE',
}

export enum OnboardingPlanStatus {
  Done = 'DONE',
  DoneLate = 'DONE_LATE',
  Late = 'LATE',
  NotStarted = 'NOT_STARTED',
  NotStartedLate = 'NOT_STARTED_LATE',
  OnTrack = 'ON_TRACK',
}

export enum OnboardingStatus {
  Done = 'DONE',
  Late = 'LATE',
  NotApplicable = 'NOT_APPLICABLE',
  NotStarted = 'NOT_STARTED',
  OnTrack = 'ON_TRACK',
  Stuck = 'STUCK',
  Successful = 'SUCCESSFUL',
}

export type OnboardingStatusInput = {
  comments?: InputMaybe<Scalars['String']['input']>;
  organizationId: Scalars['ID']['input'];
  status: OnboardingStatus;
};

export type Opportunity = Node & {
  __typename?: 'Opportunity';
  amount: Scalars['Float']['output'];
  appSource: Scalars['String']['output'];
  comments: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  estimatedClosedAt?: Maybe<Scalars['Time']['output']>;
  externalLinks: Array<ExternalSystem>;
  externalStage: Scalars['String']['output'];
  externalType: Scalars['String']['output'];
  generalNotes: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  internalStage: InternalStage;
  internalType: InternalType;
  maxAmount: Scalars['Float']['output'];
  name: Scalars['String']['output'];
  nextSteps: Scalars['String']['output'];
  owner?: Maybe<User>;
  renewalLikelihood: OpportunityRenewalLikelihood;
  renewalUpdatedByUserAt: Scalars['Time']['output'];
  renewalUpdatedByUserId: Scalars['String']['output'];
  renewedAt: Scalars['Time']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
  updatedAt: Scalars['Time']['output'];
};

export enum OpportunityRenewalLikelihood {
  HighRenewal = 'HIGH_RENEWAL',
  LowRenewal = 'LOW_RENEWAL',
  MediumRenewal = 'MEDIUM_RENEWAL',
  ZeroRenewal = 'ZERO_RENEWAL',
}

export type OpportunityRenewalUpdateInput = {
  amount?: InputMaybe<Scalars['Float']['input']>;
  appSource?: InputMaybe<Scalars['String']['input']>;
  comments?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  opportunityId: Scalars['ID']['input'];
  ownerUserId?: InputMaybe<Scalars['ID']['input']>;
  renewalLikelihood?: InputMaybe<OpportunityRenewalLikelihood>;
};

export type OpportunityUpdateInput = {
  amount?: InputMaybe<Scalars['Float']['input']>;
  appSource?: InputMaybe<Scalars['String']['input']>;
  estimatedClosedDate?: InputMaybe<Scalars['Time']['input']>;
  externalReference?: InputMaybe<ExternalSystemReferenceInput>;
  externalStage?: InputMaybe<Scalars['String']['input']>;
  externalType?: InputMaybe<Scalars['String']['input']>;
  generalNotes?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  nextSteps?: InputMaybe<Scalars['String']['input']>;
  opportunityId: Scalars['ID']['input'];
};

export type OrgAccountDetails = {
  __typename?: 'OrgAccountDetails';
  onboarding?: Maybe<OnboardingDetails>;
  renewalSummary?: Maybe<RenewalSummary>;
};

export type Organization = MetadataInterface & {
  __typename?: 'Organization';
  accountDetails?: Maybe<OrgAccountDetails>;
  /** @deprecated Use metadata.appSource */
  appSource: Scalars['String']['output'];
  contacts: ContactsPage;
  contracts?: Maybe<Array<Contract>>;
  /** @deprecated Use metadata.created */
  createdAt: Scalars['Time']['output'];
  customFields: Array<CustomField>;
  customId?: Maybe<Scalars['String']['output']>;
  customerOsId: Scalars['String']['output'];
  description?: Maybe<Scalars['String']['output']>;
  domains: Array<Scalars['String']['output']>;
  emails: Array<Email>;
  employeeGrowthRate?: Maybe<Scalars['String']['output']>;
  employees?: Maybe<Scalars['Int64']['output']>;
  entityTemplate?: Maybe<EntityTemplate>;
  externalLinks: Array<ExternalSystem>;
  fieldSets: Array<FieldSet>;
  headquarters?: Maybe<Scalars['String']['output']>;
  hide: Scalars['Boolean']['output'];
  /** @deprecated Use metadata.id */
  id: Scalars['ID']['output'];
  industry?: Maybe<Scalars['String']['output']>;
  industryGroup?: Maybe<Scalars['String']['output']>;
  isCustomer?: Maybe<Scalars['Boolean']['output']>;
  /** @deprecated Use public */
  isPublic?: Maybe<Scalars['Boolean']['output']>;
  issueSummaryByStatus: Array<IssueSummaryByStatus>;
  jobRoles: Array<JobRole>;
  lastFundingAmount?: Maybe<Scalars['String']['output']>;
  lastFundingRound?: Maybe<FundingRound>;
  /** @deprecated Use lastTouchpoint */
  lastTouchPointAt?: Maybe<Scalars['Time']['output']>;
  /** @deprecated Use lastTouchpoint */
  lastTouchPointTimelineEvent?: Maybe<TimelineEvent>;
  lastTouchPointTimelineEventId?: Maybe<Scalars['ID']['output']>;
  /** @deprecated Use lastTouchpoint */
  lastTouchPointType?: Maybe<LastTouchpointType>;
  lastTouchpoint?: Maybe<LastTouchpoint>;
  locations: Array<Location>;
  logo?: Maybe<Scalars['String']['output']>;
  /** @deprecated Use logo */
  logoUrl?: Maybe<Scalars['String']['output']>;
  market?: Maybe<Market>;
  metadata: Metadata;
  name: Scalars['String']['output'];
  /** @deprecated Use notes */
  note?: Maybe<Scalars['String']['output']>;
  notes?: Maybe<Scalars['String']['output']>;
  owner?: Maybe<User>;
  parentCompanies: Array<LinkedOrganization>;
  phoneNumbers: Array<PhoneNumber>;
  public?: Maybe<Scalars['Boolean']['output']>;
  /** @deprecated Use customId */
  referenceId?: Maybe<Scalars['String']['output']>;
  slackChannelId?: Maybe<Scalars['String']['output']>;
  socialMedia: Array<Social>;
  /** @deprecated Use socialMedia */
  socials: Array<Social>;
  /** @deprecated Use metadata.source */
  source: DataSource;
  /** @deprecated Use metadata.sourceOfTruth */
  sourceOfTruth: DataSource;
  subIndustry?: Maybe<Scalars['String']['output']>;
  subsidiaries: Array<LinkedOrganization>;
  /** @deprecated Use parentCompany */
  subsidiaryOf: Array<LinkedOrganization>;
  suggestedMergeTo: Array<SuggestedMergeOrganization>;
  tags?: Maybe<Array<Tag>>;
  targetAudience?: Maybe<Scalars['String']['output']>;
  timelineEvents: Array<TimelineEvent>;
  timelineEventsTotalCount: Scalars['Int64']['output'];
  /** @deprecated Use metadata.lastUpdated */
  updatedAt: Scalars['Time']['output'];
  valueProposition?: Maybe<Scalars['String']['output']>;
  website?: Maybe<Scalars['String']['output']>;
  yearFounded?: Maybe<Scalars['Int64']['output']>;
};

export type OrganizationContactsArgs = {
  pagination?: InputMaybe<Pagination>;
  sort?: InputMaybe<Array<SortBy>>;
  where?: InputMaybe<Filter>;
};

export type OrganizationTimelineEventsArgs = {
  from?: InputMaybe<Scalars['Time']['input']>;
  size: Scalars['Int']['input'];
  timelineEventTypes?: InputMaybe<Array<TimelineEventType>>;
};

export type OrganizationTimelineEventsTotalCountArgs = {
  timelineEventTypes?: InputMaybe<Array<TimelineEventType>>;
};

export type OrganizationInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  customFields?: InputMaybe<Array<CustomFieldInput>>;
  /**
   * The name of the organization.
   * **Required.**
   */
  customId?: InputMaybe<Scalars['String']['input']>;
  description?: InputMaybe<Scalars['String']['input']>;
  domains?: InputMaybe<Array<Scalars['String']['input']>>;
  employeeGrowthRate?: InputMaybe<Scalars['String']['input']>;
  employees?: InputMaybe<Scalars['Int64']['input']>;
  fieldSets?: InputMaybe<Array<FieldSetInput>>;
  headquarters?: InputMaybe<Scalars['String']['input']>;
  industry?: InputMaybe<Scalars['String']['input']>;
  industryGroup?: InputMaybe<Scalars['String']['input']>;
  isCustomer?: InputMaybe<Scalars['Boolean']['input']>;
  isPublic?: InputMaybe<Scalars['Boolean']['input']>;
  logo?: InputMaybe<Scalars['String']['input']>;
  logoUrl?: InputMaybe<Scalars['String']['input']>;
  market?: InputMaybe<Market>;
  name?: InputMaybe<Scalars['String']['input']>;
  note?: InputMaybe<Scalars['String']['input']>;
  notes?: InputMaybe<Scalars['String']['input']>;
  public?: InputMaybe<Scalars['Boolean']['input']>;
  referenceId?: InputMaybe<Scalars['String']['input']>;
  slackChannelId?: InputMaybe<Scalars['String']['input']>;
  subIndustry?: InputMaybe<Scalars['String']['input']>;
  templateId?: InputMaybe<Scalars['ID']['input']>;
  website?: InputMaybe<Scalars['String']['input']>;
  yearFounded?: InputMaybe<Scalars['Int64']['input']>;
};

export type OrganizationPage = Pages & {
  __typename?: 'OrganizationPage';
  content: Array<Organization>;
  totalAvailable: Scalars['Int64']['output'];
  totalElements: Scalars['Int64']['output'];
  totalPages: Scalars['Int']['output'];
};

export type OrganizationParticipant = {
  __typename?: 'OrganizationParticipant';
  organizationParticipant: Organization;
  type?: Maybe<Scalars['String']['output']>;
};

export type OrganizationPlan = Node &
  SourceFields & {
    __typename?: 'OrganizationPlan';
    appSource: Scalars['String']['output'];
    createdAt: Scalars['Time']['output'];
    id: Scalars['ID']['output'];
    masterPlanId: Scalars['ID']['output'];
    milestones: Array<OrganizationPlanMilestone>;
    name: Scalars['String']['output'];
    retired: Scalars['Boolean']['output'];
    retiredMilestones: Array<OrganizationPlanMilestone>;
    source: DataSource;
    sourceOfTruth: DataSource;
    statusDetails: OrganizationPlanStatusDetails;
    updatedAt: Scalars['Time']['output'];
  };

export type OrganizationPlanInput = {
  masterPlanId?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  organizationId: Scalars['ID']['input'];
};

export type OrganizationPlanMilestone = Node &
  SourceFields & {
    __typename?: 'OrganizationPlanMilestone';
    adhoc: Scalars['Boolean']['output'];
    appSource: Scalars['String']['output'];
    createdAt: Scalars['Time']['output'];
    dueDate: Scalars['Time']['output'];
    id: Scalars['ID']['output'];
    items: Array<OrganizationPlanMilestoneItem>;
    name: Scalars['String']['output'];
    optional: Scalars['Boolean']['output'];
    order: Scalars['Int64']['output'];
    retired: Scalars['Boolean']['output'];
    source: DataSource;
    sourceOfTruth: DataSource;
    statusDetails: OrganizationPlanMilestoneStatusDetails;
    updatedAt: Scalars['Time']['output'];
  };

export type OrganizationPlanMilestoneInput = {
  adhoc: Scalars['Boolean']['input'];
  createdAt: Scalars['Time']['input'];
  dueDate: Scalars['Time']['input'];
  items: Array<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  optional: Scalars['Boolean']['input'];
  order: Scalars['Int64']['input'];
  organizationId: Scalars['ID']['input'];
  organizationPlanId: Scalars['ID']['input'];
};

export type OrganizationPlanMilestoneItem = {
  __typename?: 'OrganizationPlanMilestoneItem';
  status: OnboardingPlanMilestoneItemStatus;
  text: Scalars['String']['output'];
  updatedAt: Scalars['Time']['output'];
  uuid: Scalars['ID']['output'];
};

export type OrganizationPlanMilestoneItemInput = {
  status: OnboardingPlanMilestoneItemStatus;
  text: Scalars['String']['input'];
  updatedAt: Scalars['Time']['input'];
  uuid?: InputMaybe<Scalars['ID']['input']>;
};

export type OrganizationPlanMilestoneReorderInput = {
  orderedIds: Array<Scalars['ID']['input']>;
  organizationId: Scalars['ID']['input'];
  organizationPlanId: Scalars['ID']['input'];
};

export type OrganizationPlanMilestoneStatusDetails = {
  __typename?: 'OrganizationPlanMilestoneStatusDetails';
  status: OnboardingPlanMilestoneStatus;
  text: Scalars['String']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type OrganizationPlanMilestoneStatusDetailsInput = {
  status: OnboardingPlanMilestoneStatus;
  text: Scalars['String']['input'];
  updatedAt: Scalars['Time']['input'];
};

export type OrganizationPlanMilestoneUpdateInput = {
  adhoc?: InputMaybe<Scalars['Boolean']['input']>;
  dueDate?: InputMaybe<Scalars['Time']['input']>;
  id: Scalars['ID']['input'];
  items?: InputMaybe<Array<InputMaybe<OrganizationPlanMilestoneItemInput>>>;
  name?: InputMaybe<Scalars['String']['input']>;
  optional?: InputMaybe<Scalars['Boolean']['input']>;
  order?: InputMaybe<Scalars['Int64']['input']>;
  organizationId: Scalars['ID']['input'];
  organizationPlanId: Scalars['ID']['input'];
  retired?: InputMaybe<Scalars['Boolean']['input']>;
  statusDetails?: InputMaybe<OrganizationPlanMilestoneStatusDetailsInput>;
  updatedAt: Scalars['Time']['input'];
};

export type OrganizationPlanStatusDetails = {
  __typename?: 'OrganizationPlanStatusDetails';
  status: OnboardingPlanStatus;
  text: Scalars['String']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type OrganizationPlanStatusDetailsInput = {
  status: OnboardingPlanStatus;
  text: Scalars['String']['input'];
  updatedAt: Scalars['Time']['input'];
};

export type OrganizationPlanUpdateInput = {
  id: Scalars['ID']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  organizationId: Scalars['ID']['input'];
  retired?: InputMaybe<Scalars['Boolean']['input']>;
  statusDetails?: InputMaybe<OrganizationPlanStatusDetailsInput>;
};

export type OrganizationUpdateInput = {
  customId?: InputMaybe<Scalars['String']['input']>;
  description?: InputMaybe<Scalars['String']['input']>;
  domains?: InputMaybe<Array<Scalars['String']['input']>>;
  employeeGrowthRate?: InputMaybe<Scalars['String']['input']>;
  employees?: InputMaybe<Scalars['Int64']['input']>;
  headquarters?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['ID']['input'];
  industry?: InputMaybe<Scalars['String']['input']>;
  industryGroup?: InputMaybe<Scalars['String']['input']>;
  isCustomer?: InputMaybe<Scalars['Boolean']['input']>;
  isPublic?: InputMaybe<Scalars['Boolean']['input']>;
  lastFundingAmount?: InputMaybe<Scalars['String']['input']>;
  lastFundingRound?: InputMaybe<FundingRound>;
  logo?: InputMaybe<Scalars['String']['input']>;
  logoUrl?: InputMaybe<Scalars['String']['input']>;
  market?: InputMaybe<Market>;
  name?: InputMaybe<Scalars['String']['input']>;
  note?: InputMaybe<Scalars['String']['input']>;
  notes?: InputMaybe<Scalars['String']['input']>;
  /** Set to true when partial update is needed. Empty or missing fields will not be ignored. */
  patch?: InputMaybe<Scalars['Boolean']['input']>;
  public?: InputMaybe<Scalars['Boolean']['input']>;
  referenceId?: InputMaybe<Scalars['String']['input']>;
  slackChannelId?: InputMaybe<Scalars['String']['input']>;
  subIndustry?: InputMaybe<Scalars['String']['input']>;
  targetAudience?: InputMaybe<Scalars['String']['input']>;
  valueProposition?: InputMaybe<Scalars['String']['input']>;
  website?: InputMaybe<Scalars['String']['input']>;
  yearFounded?: InputMaybe<Scalars['Int64']['input']>;
};

export type PageView = Node &
  SourceFields & {
    __typename?: 'PageView';
    appSource: Scalars['String']['output'];
    application: Scalars['String']['output'];
    endedAt: Scalars['Time']['output'];
    engagedTime: Scalars['Int64']['output'];
    id: Scalars['ID']['output'];
    orderInSession: Scalars['Int64']['output'];
    pageTitle: Scalars['String']['output'];
    pageUrl: Scalars['String']['output'];
    sessionId: Scalars['ID']['output'];
    source: DataSource;
    sourceOfTruth: DataSource;
    startedAt: Scalars['Time']['output'];
  };

/**
 * Describes the number of pages and total elements included in a query response.
 * **A `response` object.**
 */
export type Pages = {
  /**
   * The total number of elements included in the query response.
   * **Required.**
   */
  totalElements: Scalars['Int64']['output'];
  /**
   * The total number of pages included in the query response.
   * **Required.**
   */
  totalPages: Scalars['Int']['output'];
};

/** If provided as part of the request, results will be filtered down to the `page` and `limit` specified. */
export type Pagination = {
  /**
   * The maximum number of results in the response.
   * **Required.**
   */
  limit: Scalars['Int']['input'];
  /**
   * The results page to return in the response.
   * **Required.**
   */
  page: Scalars['Int']['input'];
};

/**
 * The honorific title of an individual.
 * **A `response` object.**
 */
export enum PersonTitle {
  /** For the holder of a doctoral degree. */
  Dr = 'DR',
  /** For girls, unmarried women, and married women who continue to use their maiden name. */
  Miss = 'MISS',
  /** For men, regardless of marital status. */
  Mr = 'MR',
  /** For married women. */
  Mrs = 'MRS',
  /** For women, regardless of marital status, or when marital status is unknown. */
  Ms = 'MS',
}

/**
 * Describes a phone number associated with a `Contact` in customerOS.
 * **A `return` object.**
 */
export type PhoneNumber = {
  __typename?: 'PhoneNumber';
  appSource?: Maybe<Scalars['String']['output']>;
  contacts: Array<Contact>;
  country?: Maybe<Country>;
  createdAt: Scalars['Time']['output'];
  /** The phone number in e164 format.  */
  e164?: Maybe<Scalars['String']['output']>;
  /**
   * The unique ID associated with the phone number.
   * **Required**
   */
  id: Scalars['ID']['output'];
  /** Defines the type of phone number. */
  label?: Maybe<PhoneNumberLabel>;
  organizations: Array<Organization>;
  /**
   * Determines if the phone number is primary or not.
   * **Required**
   */
  primary: Scalars['Boolean']['output'];
  rawPhoneNumber?: Maybe<Scalars['String']['output']>;
  source: DataSource;
  updatedAt: Scalars['Time']['output'];
  users: Array<User>;
  validated?: Maybe<Scalars['Boolean']['output']>;
};

/**
 * Describes a phone number associated with a `Contact` in customerOS.
 * **A `create` object.**
 */
export type PhoneNumberInput = {
  countryCodeA2?: InputMaybe<Scalars['String']['input']>;
  /** Defines the type of phone number. */
  label?: InputMaybe<PhoneNumberLabel>;
  /**
   * The phone number in e164 format.
   * **Required**
   */
  phoneNumber: Scalars['String']['input'];
  /**
   * Determines if the phone number is primary or not.
   * **Required**
   */
  primary?: InputMaybe<Scalars['Boolean']['input']>;
};

/**
 * Defines the type of phone number.
 * **A `response` object. **
 */
export enum PhoneNumberLabel {
  Home = 'HOME',
  Main = 'MAIN',
  Mobile = 'MOBILE',
  Other = 'OTHER',
  Work = 'WORK',
}

export type PhoneNumberParticipant = {
  __typename?: 'PhoneNumberParticipant';
  phoneNumberParticipant: PhoneNumber;
  type?: Maybe<Scalars['String']['output']>;
};

/**
 * Describes a phone number associated with a `Contact` in customerOS.
 * **An `update` object.**
 */
export type PhoneNumberUpdateInput = {
  countryCodeA2?: InputMaybe<Scalars['String']['input']>;
  /**
   * The unique ID associated with the phone number.
   * **Required**
   */
  id: Scalars['ID']['input'];
  /** Defines the type of phone number. */
  label?: InputMaybe<PhoneNumberLabel>;
  phoneNumber?: InputMaybe<Scalars['String']['input']>;
  /**
   * Determines if the phone number is primary or not.
   * **Required**
   */
  primary?: InputMaybe<Scalars['Boolean']['input']>;
};

export type Player = {
  __typename?: 'Player';
  appSource: Scalars['String']['output'];
  authId: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  identityId?: Maybe<Scalars['String']['output']>;
  provider: Scalars['String']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
  updatedAt: Scalars['Time']['output'];
  users: Array<PlayerUser>;
};

export type PlayerInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  authId: Scalars['String']['input'];
  identityId?: InputMaybe<Scalars['String']['input']>;
  provider: Scalars['String']['input'];
};

export type PlayerUpdate = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  identityId?: InputMaybe<Scalars['String']['input']>;
};

export type PlayerUser = {
  __typename?: 'PlayerUser';
  default: Scalars['Boolean']['output'];
  tenant: Scalars['String']['output'];
  user: User;
};

export type Query = {
  __typename?: 'Query';
  analysis: Analysis;
  attachment: Attachment;
  bankAccounts: Array<BankAccount>;
  billableInfo: TenantBillableInfo;
  /** Fetch a single contact from customerOS by contact ID. */
  contact?: Maybe<Contact>;
  contact_ByEmail: Contact;
  contact_ByPhone: Contact;
  /**
   * Fetch paginated list of contacts
   * Possible values for sort:
   * - PREFIX
   * - FIRST_NAME
   * - LAST_NAME
   * - NAME
   * - DESCRIPTION
   * - CREATED_AT
   */
  contacts: ContactsPage;
  contract: Contract;
  /** sort.By available options: ORGANIZATION, IS_CUSTOMER, DOMAIN, LOCATION, OWNER, LAST_TOUCHPOINT, RENEWAL_LIKELIHOOD, FORECAST_ARR, RENEWAL_DATE, ONBOARDING_STATUS */
  dashboardView_Organizations?: Maybe<OrganizationPage>;
  dashboardView_Renewals?: Maybe<RenewalsPage>;
  dashboard_ARRBreakdown?: Maybe<DashboardArrBreakdown>;
  dashboard_CustomerMap?: Maybe<Array<DashboardCustomerMap>>;
  dashboard_GrossRevenueRetention?: Maybe<DashboardGrossRevenueRetention>;
  dashboard_MRRPerCustomer?: Maybe<DashboardMrrPerCustomer>;
  dashboard_NewCustomers?: Maybe<DashboardNewCustomers>;
  dashboard_OnboardingCompletion?: Maybe<DashboardOnboardingCompletion>;
  dashboard_RetentionRate?: Maybe<DashboardRetentionRate>;
  dashboard_RevenueAtRisk?: Maybe<DashboardRevenueAtRisk>;
  dashboard_TimeToOnboard?: Maybe<DashboardTimeToOnboard>;
  email: Email;
  entityTemplates: Array<EntityTemplate>;
  externalMeetings: MeetingsPage;
  gcli_Search: Array<GCliItem>;
  global_Cache: GlobalCache;
  interactionEvent: InteractionEvent;
  interactionEvent_ByEventIdentifier: InteractionEvent;
  interactionSession: InteractionSession;
  interactionSession_BySessionIdentifier: InteractionSession;
  invoice: Invoice;
  invoices: InvoicesPage;
  invoicingCycle: InvoicingCycle;
  issue: Issue;
  logEntry: LogEntry;
  masterPlan: MasterPlan;
  masterPlans: Array<MasterPlan>;
  meeting: Meeting;
  opportunity?: Maybe<Opportunity>;
  organization?: Maybe<Organization>;
  organizationPlan: OrganizationPlan;
  organizationPlans: Array<OrganizationPlan>;
  organizationPlansForOrganization: Array<OrganizationPlan>;
  organization_ByCustomerOsId?: Maybe<Organization>;
  organization_DistinctOwners: Array<User>;
  organizations: OrganizationPage;
  phoneNumber: PhoneNumber;
  player_ByAuthIdProvider: Player;
  reminder: Reminder;
  remindersForOrganization: Array<Reminder>;
  serviceLineItem: ServiceLineItem;
  slack_Channels: SlackChannelPage;
  tableViewDefs: TableViewDefPage;
  tags: Array<Tag>;
  tenant: Scalars['String']['output'];
  tenantBillingProfile: TenantBillingProfile;
  tenantBillingProfiles: Array<TenantBillingProfile>;
  tenantSettings: TenantSettings;
  tenant_ByEmail?: Maybe<Scalars['String']['output']>;
  tenant_ByWorkspace?: Maybe<Scalars['String']['output']>;
  timelineEvents: Array<TimelineEvent>;
  user: User;
  user_ByEmail: User;
  users: UserPage;
};

export type QueryAnalysisArgs = {
  id: Scalars['ID']['input'];
};

export type QueryAttachmentArgs = {
  id: Scalars['ID']['input'];
};

export type QueryContactArgs = {
  id: Scalars['ID']['input'];
};

export type QueryContact_ByEmailArgs = {
  email: Scalars['String']['input'];
};

export type QueryContact_ByPhoneArgs = {
  e164: Scalars['String']['input'];
};

export type QueryContactsArgs = {
  pagination?: InputMaybe<Pagination>;
  sort?: InputMaybe<Array<SortBy>>;
  where?: InputMaybe<Filter>;
};

export type QueryContractArgs = {
  id: Scalars['ID']['input'];
};

export type QueryDashboardView_OrganizationsArgs = {
  pagination: Pagination;
  sort?: InputMaybe<SortBy>;
  where?: InputMaybe<Filter>;
};

export type QueryDashboardView_RenewalsArgs = {
  pagination: Pagination;
  sort?: InputMaybe<SortBy>;
  where?: InputMaybe<Filter>;
};

export type QueryDashboard_ArrBreakdownArgs = {
  period?: InputMaybe<DashboardPeriodInput>;
};

export type QueryDashboard_GrossRevenueRetentionArgs = {
  period?: InputMaybe<DashboardPeriodInput>;
};

export type QueryDashboard_MrrPerCustomerArgs = {
  period?: InputMaybe<DashboardPeriodInput>;
};

export type QueryDashboard_NewCustomersArgs = {
  period?: InputMaybe<DashboardPeriodInput>;
};

export type QueryDashboard_OnboardingCompletionArgs = {
  period?: InputMaybe<DashboardPeriodInput>;
};

export type QueryDashboard_RetentionRateArgs = {
  period?: InputMaybe<DashboardPeriodInput>;
};

export type QueryDashboard_RevenueAtRiskArgs = {
  period?: InputMaybe<DashboardPeriodInput>;
};

export type QueryDashboard_TimeToOnboardArgs = {
  period?: InputMaybe<DashboardPeriodInput>;
};

export type QueryEmailArgs = {
  id: Scalars['ID']['input'];
};

export type QueryEntityTemplatesArgs = {
  extends?: InputMaybe<EntityTemplateExtension>;
};

export type QueryExternalMeetingsArgs = {
  externalId?: InputMaybe<Scalars['ID']['input']>;
  externalSystemId: Scalars['String']['input'];
  pagination?: InputMaybe<Pagination>;
  sort?: InputMaybe<Array<SortBy>>;
  where?: InputMaybe<Filter>;
};

export type QueryGcli_SearchArgs = {
  keyword: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
};

export type QueryInteractionEventArgs = {
  id: Scalars['ID']['input'];
};

export type QueryInteractionEvent_ByEventIdentifierArgs = {
  eventIdentifier: Scalars['String']['input'];
};

export type QueryInteractionSessionArgs = {
  id: Scalars['ID']['input'];
};

export type QueryInteractionSession_BySessionIdentifierArgs = {
  sessionIdentifier: Scalars['String']['input'];
};

export type QueryInvoiceArgs = {
  id: Scalars['ID']['input'];
};

export type QueryInvoicesArgs = {
  organizationId?: InputMaybe<Scalars['ID']['input']>;
  pagination?: InputMaybe<Pagination>;
  sort?: InputMaybe<Array<SortBy>>;
  where?: InputMaybe<Filter>;
};

export type QueryIssueArgs = {
  id: Scalars['ID']['input'];
};

export type QueryLogEntryArgs = {
  id: Scalars['ID']['input'];
};

export type QueryMasterPlanArgs = {
  id: Scalars['ID']['input'];
};

export type QueryMasterPlansArgs = {
  retired?: InputMaybe<Scalars['Boolean']['input']>;
};

export type QueryMeetingArgs = {
  id: Scalars['ID']['input'];
};

export type QueryOpportunityArgs = {
  id: Scalars['ID']['input'];
};

export type QueryOrganizationArgs = {
  id: Scalars['ID']['input'];
};

export type QueryOrganizationPlanArgs = {
  id: Scalars['ID']['input'];
};

export type QueryOrganizationPlansArgs = {
  retired?: InputMaybe<Scalars['Boolean']['input']>;
};

export type QueryOrganizationPlansForOrganizationArgs = {
  organizationId: Scalars['ID']['input'];
};

export type QueryOrganization_ByCustomerOsIdArgs = {
  customerOsId: Scalars['String']['input'];
};

export type QueryOrganizationsArgs = {
  pagination?: InputMaybe<Pagination>;
  sort?: InputMaybe<Array<SortBy>>;
  where?: InputMaybe<Filter>;
};

export type QueryPhoneNumberArgs = {
  id: Scalars['ID']['input'];
};

export type QueryPlayer_ByAuthIdProviderArgs = {
  authId: Scalars['String']['input'];
  provider: Scalars['String']['input'];
};

export type QueryReminderArgs = {
  id: Scalars['ID']['input'];
};

export type QueryRemindersForOrganizationArgs = {
  organizationId: Scalars['ID']['input'];
};

export type QueryServiceLineItemArgs = {
  id: Scalars['ID']['input'];
};

export type QuerySlack_ChannelsArgs = {
  pagination?: InputMaybe<Pagination>;
};

export type QueryTableViewDefsArgs = {
  pagination?: InputMaybe<Pagination>;
  sort?: InputMaybe<SortBy>;
  where?: InputMaybe<Filter>;
};

export type QueryTenantBillingProfileArgs = {
  id: Scalars['ID']['input'];
};

export type QueryTenant_ByEmailArgs = {
  email: Scalars['String']['input'];
};

export type QueryTenant_ByWorkspaceArgs = {
  workspace: WorkspaceInput;
};

export type QueryTimelineEventsArgs = {
  ids: Array<Scalars['ID']['input']>;
};

export type QueryUserArgs = {
  id: Scalars['ID']['input'];
};

export type QueryUser_ByEmailArgs = {
  email: Scalars['String']['input'];
};

export type QueryUsersArgs = {
  pagination?: InputMaybe<Pagination>;
  sort?: InputMaybe<Array<SortBy>>;
  where?: InputMaybe<Filter>;
};

export type Reminder = MetadataInterface & {
  __typename?: 'Reminder';
  content: Scalars['String']['output'];
  dismissed: Scalars['Boolean']['output'];
  dueDate: Scalars['Time']['output'];
  metadata: Metadata;
  owner: User;
};

export type ReminderInput = {
  content: Scalars['String']['input'];
  dueDate: Scalars['Time']['input'];
  organizationId: Scalars['ID']['input'];
  userId: Scalars['ID']['input'];
};

export type ReminderUpdateInput = {
  content?: InputMaybe<Scalars['String']['input']>;
  dismissed?: InputMaybe<Scalars['Boolean']['input']>;
  dueDate?: InputMaybe<Scalars['Time']['input']>;
  id: Scalars['ID']['input'];
};

export type RenewalRecord = {
  __typename?: 'RenewalRecord';
  contract: Contract;
  opportunity?: Maybe<Opportunity>;
  organization: Organization;
};

export type RenewalSummary = {
  __typename?: 'RenewalSummary';
  arrForecast?: Maybe<Scalars['Float']['output']>;
  maxArrForecast?: Maybe<Scalars['Float']['output']>;
  nextRenewalDate?: Maybe<Scalars['Time']['output']>;
  renewalLikelihood?: Maybe<OpportunityRenewalLikelihood>;
};

export type RenewalsPage = Pages & {
  __typename?: 'RenewalsPage';
  content: Array<RenewalRecord>;
  totalAvailable: Scalars['Int64']['output'];
  totalElements: Scalars['Int64']['output'];
  totalPages: Scalars['Int']['output'];
};

/**
 * Describes the success or failure of the GraphQL call.
 * **A `return` object**
 */
export type Result = {
  __typename?: 'Result';
  /**
   * The result of the GraphQL call.
   * **Required.**
   */
  result: Scalars['Boolean']['output'];
};

export enum Role {
  Admin = 'ADMIN',
  CustomerOsPlatformOwner = 'CUSTOMER_OS_PLATFORM_OWNER',
  Owner = 'OWNER',
  User = 'USER',
}

export type ServiceLineItem = MetadataInterface & {
  __typename?: 'ServiceLineItem';
  billingCycle: BilledType;
  comments: Scalars['String']['output'];
  createdBy?: Maybe<User>;
  description: Scalars['String']['output'];
  externalLinks: Array<ExternalSystem>;
  metadata: Metadata;
  parentId: Scalars['ID']['output'];
  price: Scalars['Float']['output'];
  quantity: Scalars['Int64']['output'];
  serviceEnded?: Maybe<Scalars['Time']['output']>;
  serviceStarted: Scalars['Time']['output'];
  tax: Tax;
};

export type ServiceLineItemBulkUpdateInput = {
  contractId: Scalars['ID']['input'];
  invoiceNote?: InputMaybe<Scalars['String']['input']>;
  serviceLineItems: Array<InputMaybe<ServiceLineItemBulkUpdateItem>>;
};

export type ServiceLineItemBulkUpdateItem = {
  billed?: InputMaybe<BilledType>;
  comments?: InputMaybe<Scalars['String']['input']>;
  isRetroactiveCorrection?: InputMaybe<Scalars['Boolean']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  price?: InputMaybe<Scalars['Float']['input']>;
  quantity?: InputMaybe<Scalars['Int64']['input']>;
  serviceLineItemId?: InputMaybe<Scalars['ID']['input']>;
  serviceStarted?: InputMaybe<Scalars['Time']['input']>;
  vatRate?: InputMaybe<Scalars['Float']['input']>;
};

export type ServiceLineItemCloseInput = {
  endedAt?: InputMaybe<Scalars['Time']['input']>;
  id: Scalars['ID']['input'];
  serviceEnded?: InputMaybe<Scalars['Time']['input']>;
};

export type ServiceLineItemInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  billingCycle?: InputMaybe<BilledType>;
  contractId: Scalars['ID']['input'];
  description?: InputMaybe<Scalars['String']['input']>;
  price?: InputMaybe<Scalars['Float']['input']>;
  quantity?: InputMaybe<Scalars['Int64']['input']>;
  serviceEnded?: InputMaybe<Scalars['Time']['input']>;
  serviceStarted?: InputMaybe<Scalars['Time']['input']>;
  tax?: InputMaybe<TaxInput>;
};

export type ServiceLineItemUpdateInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  billingCycle?: InputMaybe<BilledType>;
  comments?: InputMaybe<Scalars['String']['input']>;
  description?: InputMaybe<Scalars['String']['input']>;
  id?: InputMaybe<Scalars['ID']['input']>;
  isRetroactiveCorrection?: InputMaybe<Scalars['Boolean']['input']>;
  price?: InputMaybe<Scalars['Float']['input']>;
  quantity?: InputMaybe<Scalars['Int64']['input']>;
  serviceEnded?: InputMaybe<Scalars['Time']['input']>;
  serviceStarted?: InputMaybe<Scalars['Time']['input']>;
  tax?: InputMaybe<TaxInput>;
};

export type SlackChannel = {
  __typename?: 'SlackChannel';
  channelId: Scalars['String']['output'];
  channelName: Scalars['String']['output'];
  metadata: Metadata;
  organization?: Maybe<Organization>;
};

export type SlackChannelPage = Pages & {
  __typename?: 'SlackChannelPage';
  content: Array<SlackChannel>;
  totalAvailable: Scalars['Int64']['output'];
  totalElements: Scalars['Int64']['output'];
  totalPages: Scalars['Int']['output'];
};

export type Social = Node &
  SourceFields & {
    __typename?: 'Social';
    appSource: Scalars['String']['output'];
    createdAt: Scalars['Time']['output'];
    id: Scalars['ID']['output'];
    platformName?: Maybe<Scalars['String']['output']>;
    source: DataSource;
    sourceOfTruth: DataSource;
    updatedAt: Scalars['Time']['output'];
    url: Scalars['String']['output'];
  };

export type SocialInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  platformName?: InputMaybe<Scalars['String']['input']>;
  url: Scalars['String']['input'];
};

export type SocialUpdateInput = {
  id: Scalars['ID']['input'];
  platformName?: InputMaybe<Scalars['String']['input']>;
  url: Scalars['String']['input'];
};

export type SortBy = {
  by: Scalars['String']['input'];
  caseSensitive?: InputMaybe<Scalars['Boolean']['input']>;
  direction?: SortingDirection;
};

export enum SortingDirection {
  Asc = 'ASC',
  Desc = 'DESC',
}

export type SourceFields = {
  appSource: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
};

export type SourceFieldsInterface = {
  appSource: Scalars['String']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
};

export type State = {
  __typename?: 'State';
  code: Scalars['String']['output'];
  country: Country;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type SuggestedMergeOrganization = {
  __typename?: 'SuggestedMergeOrganization';
  confidence?: Maybe<Scalars['Float']['output']>;
  organization: Organization;
  suggestedAt?: Maybe<Scalars['Time']['output']>;
  suggestedBy?: Maybe<Scalars['String']['output']>;
};

export type TableViewDef = Node & {
  __typename?: 'TableViewDef';
  columns?: Maybe<Array<Maybe<ColumnDef>>>;
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  filters?: Maybe<Scalars['String']['output']>;
  icon?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  order?: Maybe<Scalars['Int']['output']>;
  sorting?: Maybe<Scalars['String']['output']>;
  type?: Maybe<ViewType>;
  updatedAt: Scalars['Time']['output'];
};

export type TableViewDefPage = Pages & {
  __typename?: 'TableViewDefPage';
  content: Array<TableViewDef>;
  totalAvailable: Scalars['Int64']['output'];
  totalElements: Scalars['Int64']['output'];
  totalPages: Scalars['Int']['output'];
};

export type Tag = {
  __typename?: 'Tag';
  appSource: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  source: DataSource;
  updatedAt: Scalars['Time']['output'];
};

export type TagIdOrNameInput = {
  id?: InputMaybe<Scalars['ID']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
};

export type TagInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
};

export type TagUpdateInput = {
  id: Scalars['ID']['input'];
  name: Scalars['String']['input'];
};

export type Tax = {
  __typename?: 'Tax';
  salesTax: Scalars['Boolean']['output'];
  taxRate: Scalars['Float']['output'];
  vat: Scalars['Boolean']['output'];
};

export type TaxInput = {
  taxRate: Scalars['Float']['input'];
};

export type TenantBillableInfo = {
  __typename?: 'TenantBillableInfo';
  greylistedContacts: Scalars['Int64']['output'];
  greylistedOrganizations: Scalars['Int64']['output'];
  whitelistedContacts: Scalars['Int64']['output'];
  whitelistedOrganizations: Scalars['Int64']['output'];
};

export type TenantBillingProfile = Node &
  SourceFields & {
    __typename?: 'TenantBillingProfile';
    addressLine1: Scalars['String']['output'];
    addressLine2: Scalars['String']['output'];
    addressLine3: Scalars['String']['output'];
    appSource: Scalars['String']['output'];
    canPayWithBankTransfer: Scalars['Boolean']['output'];
    /** @deprecated Not used */
    canPayWithCard: Scalars['Boolean']['output'];
    /** @deprecated Not used */
    canPayWithDirectDebitACH: Scalars['Boolean']['output'];
    /** @deprecated Not used */
    canPayWithDirectDebitBacs: Scalars['Boolean']['output'];
    /** @deprecated Not used */
    canPayWithDirectDebitSEPA: Scalars['Boolean']['output'];
    canPayWithPigeon: Scalars['Boolean']['output'];
    country: Scalars['String']['output'];
    createdAt: Scalars['Time']['output'];
    domesticPaymentsBankInfo: Scalars['String']['output'];
    /** @deprecated Use sendInvoicesFrom */
    email: Scalars['String']['output'];
    id: Scalars['ID']['output'];
    internationalPaymentsBankInfo: Scalars['String']['output'];
    legalName: Scalars['String']['output'];
    locality: Scalars['String']['output'];
    phone: Scalars['String']['output'];
    sendInvoicesBcc: Scalars['String']['output'];
    sendInvoicesFrom: Scalars['String']['output'];
    source: DataSource;
    sourceOfTruth: DataSource;
    updatedAt: Scalars['Time']['output'];
    vatNumber: Scalars['String']['output'];
    zip: Scalars['String']['output'];
  };

export type TenantBillingProfileInput = {
  addressLine1?: InputMaybe<Scalars['String']['input']>;
  addressLine2?: InputMaybe<Scalars['String']['input']>;
  addressLine3?: InputMaybe<Scalars['String']['input']>;
  canPayWithBankTransfer: Scalars['Boolean']['input'];
  canPayWithCard: Scalars['Boolean']['input'];
  canPayWithDirectDebitACH: Scalars['Boolean']['input'];
  canPayWithDirectDebitBacs: Scalars['Boolean']['input'];
  canPayWithDirectDebitSEPA: Scalars['Boolean']['input'];
  canPayWithPigeon: Scalars['Boolean']['input'];
  country?: InputMaybe<Scalars['String']['input']>;
  domesticPaymentsBankInfo?: InputMaybe<Scalars['String']['input']>;
  email?: InputMaybe<Scalars['String']['input']>;
  internationalPaymentsBankInfo?: InputMaybe<Scalars['String']['input']>;
  legalName?: InputMaybe<Scalars['String']['input']>;
  locality?: InputMaybe<Scalars['String']['input']>;
  phone?: InputMaybe<Scalars['String']['input']>;
  sendInvoicesBcc?: InputMaybe<Scalars['String']['input']>;
  sendInvoicesFrom: Scalars['String']['input'];
  vatNumber: Scalars['String']['input'];
  zip?: InputMaybe<Scalars['String']['input']>;
};

export type TenantBillingProfileUpdateInput = {
  addressLine1?: InputMaybe<Scalars['String']['input']>;
  addressLine2?: InputMaybe<Scalars['String']['input']>;
  addressLine3?: InputMaybe<Scalars['String']['input']>;
  canPayWithBankTransfer?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithCard?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithDirectDebitACH?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithDirectDebitBacs?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithDirectDebitSEPA?: InputMaybe<Scalars['Boolean']['input']>;
  canPayWithPigeon?: InputMaybe<Scalars['Boolean']['input']>;
  country?: InputMaybe<Scalars['String']['input']>;
  domesticPaymentsBankInfo?: InputMaybe<Scalars['String']['input']>;
  email?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['ID']['input'];
  internationalPaymentsBankInfo?: InputMaybe<Scalars['String']['input']>;
  legalName?: InputMaybe<Scalars['String']['input']>;
  locality?: InputMaybe<Scalars['String']['input']>;
  patch?: InputMaybe<Scalars['Boolean']['input']>;
  phone?: InputMaybe<Scalars['String']['input']>;
  sendInvoicesBcc?: InputMaybe<Scalars['String']['input']>;
  sendInvoicesFrom?: InputMaybe<Scalars['String']['input']>;
  vatNumber?: InputMaybe<Scalars['String']['input']>;
  zip?: InputMaybe<Scalars['String']['input']>;
};

export type TenantInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
};

export type TenantSettings = {
  __typename?: 'TenantSettings';
  baseCurrency?: Maybe<Currency>;
  billingEnabled: Scalars['Boolean']['output'];
  logoRepositoryFileId?: Maybe<Scalars['String']['output']>;
  /** @deprecated Use logoRepositoryFileId */
  logoUrl: Scalars['String']['output'];
};

export type TenantSettingsInput = {
  baseCurrency?: InputMaybe<Currency>;
  billingEnabled?: InputMaybe<Scalars['Boolean']['input']>;
  logoRepositoryFileId?: InputMaybe<Scalars['String']['input']>;
  logoUrl?: InputMaybe<Scalars['String']['input']>;
  patch?: InputMaybe<Scalars['Boolean']['input']>;
};

export type TimeRange = {
  /**
   * The start time of the time range.
   * **Required.**
   */
  from: Scalars['Time']['input'];
  /**
   * The end time of the time range.
   * **Required.**
   */
  to: Scalars['Time']['input'];
};

export type TimelineEvent =
  | Action
  | Analysis
  | InteractionEvent
  | InteractionSession
  | Issue
  | LogEntry
  | Meeting
  | Note
  | PageView;

export enum TimelineEventType {
  Action = 'ACTION',
  Analysis = 'ANALYSIS',
  InteractionEvent = 'INTERACTION_EVENT',
  InteractionSession = 'INTERACTION_SESSION',
  Issue = 'ISSUE',
  LogEntry = 'LOG_ENTRY',
  Meeting = 'MEETING',
  Note = 'NOTE',
  PageView = 'PAGE_VIEW',
}

/**
 * Describes the User of customerOS.  A user is the person who logs into the Openline platform.
 * **A `return` object**
 */
export type User = {
  __typename?: 'User';
  appSource: Scalars['String']['output'];
  bot: Scalars['Boolean']['output'];
  calendars: Array<Calendar>;
  /**
   * Timestamp of user creation.
   * **Required**
   */
  createdAt: Scalars['Time']['output'];
  /**
   * All email addresses associated with a user in customerOS.
   * **Required.  If no values it returns an empty array.**
   */
  emails?: Maybe<Array<Email>>;
  /**
   * The first name of the customerOS user.
   * **Required**
   */
  firstName: Scalars['String']['output'];
  /**
   * The unique ID associated with the customerOS user.
   * **Required**
   */
  id: Scalars['ID']['output'];
  internal: Scalars['Boolean']['output'];
  jobRoles: Array<JobRole>;
  /**
   * The last name of the customerOS user.
   * **Required**
   */
  lastName: Scalars['String']['output'];
  name?: Maybe<Scalars['String']['output']>;
  phoneNumbers: Array<PhoneNumber>;
  player: Player;
  profilePhotoUrl?: Maybe<Scalars['String']['output']>;
  roles: Array<Role>;
  source: DataSource;
  sourceOfTruth: DataSource;
  timezone?: Maybe<Scalars['String']['output']>;
  updatedAt: Scalars['Time']['output'];
};

/**
 * Describes the User of customerOS.  A user is the person who logs into the Openline platform.
 * **A `create` object.**
 */
export type UserInput = {
  /**
   * The name of the app performing the create.
   * **Optional**
   */
  appSource?: InputMaybe<Scalars['String']['input']>;
  /**
   * The email address of the customerOS user.
   * **Required**
   */
  email: EmailInput;
  /**
   * The first name of the customerOS user.
   * **Required**
   */
  firstName: Scalars['String']['input'];
  /**
   * The Job Roles of the user.
   * **Optional**
   */
  jobRoles?: InputMaybe<Array<JobRoleInput>>;
  /**
   * The last name of the customerOS user.
   * **Required**
   */
  lastName: Scalars['String']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  /**
   * Player to associate with the user with. If the person does not exist, it will be created.
   * **Required**
   */
  player: PlayerInput;
  profilePhotoUrl?: InputMaybe<Scalars['String']['input']>;
  timezone?: InputMaybe<Scalars['String']['input']>;
};

/**
 * Specifies how many pages of `User` information has been returned in the query response.
 * **A `return` object.**
 */
export type UserPage = Pages & {
  __typename?: 'UserPage';
  /**
   * A `User` entity in customerOS.
   * **Required.  If no values it returns an empty array.**
   */
  content: Array<User>;
  /**
   * Total number of elements in the query response.
   * **Required.**
   */
  totalElements: Scalars['Int64']['output'];
  /**
   * Total number of pages in the query response.
   * **Required.**
   */
  totalPages: Scalars['Int']['output'];
};

export type UserParticipant = {
  __typename?: 'UserParticipant';
  type?: Maybe<Scalars['String']['output']>;
  userParticipant: User;
};

export type UserUpdateInput = {
  /**
   * The first name of the customerOS user.
   * **Required**
   */
  firstName: Scalars['String']['input'];
  id: Scalars['ID']['input'];
  /**
   * The last name of the customerOS user.
   * **Required**
   */
  lastName: Scalars['String']['input'];
  name?: InputMaybe<Scalars['String']['input']>;
  profilePhotoUrl?: InputMaybe<Scalars['String']['input']>;
  timezone?: InputMaybe<Scalars['String']['input']>;
};

export type ViewType = Node & {
  __typename?: 'ViewType';
  createdAt: Scalars['Time']['output'];
  createdBy?: Maybe<User>;
  id: Scalars['ID']['output'];
  name?: Maybe<Scalars['String']['output']>;
  updatedAt: Scalars['Time']['output'];
};

export type Workspace = {
  __typename?: 'Workspace';
  appSource: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  provider: Scalars['String']['output'];
  source: DataSource;
  sourceOfTruth: DataSource;
  updatedAt: Scalars['Time']['output'];
};

export type WorkspaceInput = {
  appSource?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  provider: Scalars['String']['input'];
};
