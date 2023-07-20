export interface ComposeEmailDtoI {
  to: Array<string>;
  cc: Array<string>;
  bcc: Array<string>;
  subject: string;
  content: string;
  files: Array<any>;
}

export class ComposeEmailDto implements ComposeEmailDtoI {
  to: Array<string>;
  cc: Array<string>;
  bcc: Array<string>;
  subject: string;
  content: string;
  files: Array<any>;

  constructor(data?: any) {
    this.to = data?.to || [];
    this.cc = data?.cc || [];
    this.bcc = data?.bcc || [];
    this.subject = data?.subject || '';
    this.content = data?.content || '';
    this.files = data?.files || [];
  }

  static toForm(data: any) {
    return new ComposeEmailDto(data);
  }

  static toPayload(data: ComposeEmailDtoI) {
    return {
      to: data.to,
      cc: data.cc,
      bcc: data.bcc,
      subject: data.subject,
      content: data.content,
      files: data.files,
    } as any;
  }
}
