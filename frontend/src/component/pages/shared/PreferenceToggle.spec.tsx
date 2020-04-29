import React from 'react';

import { PreferenceToggle } from './PreferenceToggle';
import { fireEvent, render } from '@testing-library/react';
import { IPagedShortLinks } from '../../../service/ShortLink.service';

describe('PreferenceToggle component', () => {
    const TOGGLE_CLASS_NAME = ".toggle";
    const TOGGLE_SELECTOR = TOGGLE_CLASS_NAME;
    const TOGGLE_ACTIVE_SELECTOR = TOGGLE_CLASS_NAME + "> .active";

    beforeAll(() => {
        jest.useFakeTimers();
    });

    test('should render without crash', () => {
        render(<PreferenceToggle onToggleClick={jest.fn} />);
    });

    test('should render the label if a label is provided', () => {
        const TEST_LABEL = "Test Toggle Preference";

        const { container } = render(
            <PreferenceToggle toggleLabel={TEST_LABEL} onToggleClick={jest.fn} />
        );

        expect(container.textContent).toContain(TEST_LABEL);
    });

    test('should render without a label if no label is provided', () => {
        const { container } = render(
            <PreferenceToggle onToggleClick={jest.fn} />
        );

        expect(container.textContent).toBeFalsy();
    });

    test('should render a toggle', () => {
        const { container } = render(
            <PreferenceToggle onToggleClick={jest.fn} />
        );

        expect(container.querySelector(TOGGLE_SELECTOR)).not.toBeNull();
    });

    test('should render an active toggle if enabled by default', () => {
        const { container } = render(
            <PreferenceToggle defaultIsEnabled={true} onToggleClick={jest.fn} />
        );

        expect(container.querySelector(TOGGLE_ACTIVE_SELECTOR)).not.toBeNull();
    });

    test('should render an inactive toggle if explicitly disabled by default', () => {
        const { container } = render(
            <PreferenceToggle defaultIsEnabled={false} onToggleClick={jest.fn} />
        );

        expect(container.querySelector(TOGGLE_SELECTOR)).not.toBeNull();
    });
});