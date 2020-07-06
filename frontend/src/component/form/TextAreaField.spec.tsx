import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { TextAreaField } from './TextAreaField';

describe('TextAreaField component', () => {
  test('should render without crash', () => {
    render(<TextAreaField />);
  });

  test('should render with correct value', () => {
    const { container } = render(<TextAreaField value={'text'} />);
    const textarea = container.querySelector('textarea');

    expect(textarea).toBeTruthy();
    expect(textarea!.value).toMatch('text');
  });

  test('should render with correct placeholder', () => {
    const { container } = render(<TextAreaField placeholder={'placeholder'} />);
    const textarea = container.querySelector('textarea');

    expect(textarea).toBeTruthy();
    expect(textarea!.placeholder).toMatch('placeholder');
  });

  test('should fire onChange when input changes', () => {
    const onChange = jest.fn();
    const { container } = render(<TextAreaField onChange={onChange} />);
    const textarea = container.querySelector('textarea');

    expect(textarea).toBeTruthy();
    fireEvent.change(textarea!, { target: { value: 'text' } });

    expect(onChange).toHaveBeenCalledTimes(1);
    expect(onChange).toHaveBeenLastCalledWith('text');
  });

  test('should fire onBlur when textarea blurs', () => {
    const onBlur = jest.fn();
    const { container } = render(<TextAreaField onBlur={onBlur} />);
    const textarea = container.querySelector('textarea');

    expect(textarea).toBeTruthy();
    fireEvent.blur(textarea!);

    expect(onBlur).toHaveBeenCalledTimes(1);
  });
});
