export function validateLongLinkFormat(longLink?: string): string | null {
    const LONG_LINK_MAX_LENGTH = 200;

    if (!longLink) {
        return 'Long link can\'t be empty';
    }

    if (longLink.length >= LONG_LINK_MAX_LENGTH) {
        return `Long link has to contain less than ${LONG_LINK_MAX_LENGTH} characters`;
    }

    if (!isUri(longLink)) {
        return 'Long link has to be in format like http://www.google.com';
    }

    return null;
}

function isUri(text: string): boolean {
    return /^[a-zA-Z]+:\/\/.+$/.test(text);
}


