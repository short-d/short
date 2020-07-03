import { EmotionType } from './emotion';

export interface Feedback {
  emotion: EmotionType;
  message?: string;
  contactEmail?: string;
}
