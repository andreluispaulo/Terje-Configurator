export interface Metadata {
  type: string;
  default: string;
  description: string;
}

export interface CFGLine {
  index: number;
  type: number; // 0=Unknown, 1=Config, 2=Comment, 3=Empty
  key: string;
  value: string;
  metadata: Metadata;
}

export interface CFGFile {
  lines: CFGLine[];
}

export interface Segment {
    isAttribute: boolean;
    content: string;
    attrName?: string;
    attrValue?: string;
}

export interface XMLLine {
    index: number;
    segments: Segment[];
    depth: number;
    tagName: string;
}

export interface XMLFile {
    lines: XMLLine[];
}

export type FileContent = CFGFile | XMLFile;
