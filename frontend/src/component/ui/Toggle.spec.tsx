import React from 'react';
import { render } from '@testing-library/react';
import { Toggle } from './Toggle';

describe('Toggle component', () => {
    test('should render without crash', () => {
        render(<Toggle />);
    });

    test('should render an inactive toggle when disabled by default', () => {
        fail("Not implemented");
    });

    test('should switch from disabled to enabled when clicked', () => {
        const toggleRef = React.createRef<Toggle>();
        const { container } = render(<Toggle ref={toggleRef} defaultIsEnabled={false} />);

        toggleRef.current?.handleClick();
        expect(container.querySelector(".background.active")).not.toBeNull();
        expect(container.querySelector(".knob.active")).not.toBeNull();
    });

    test('should trigger onClick callback when toggle clicked', () => {
        fail("Not implemented");
    });
});
