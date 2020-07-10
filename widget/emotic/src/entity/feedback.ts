import { EmotionType } from './emotion';

export interface Feedback {
  emotion: EmotionType;
  comment?: string;
  contactEmail?: string;
}
