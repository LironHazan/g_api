import { render } from '@testing-library/react';

import ReproIssue15475Remote from './repro-issue-15475-remote';

describe('ReproIssue15475Remote', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<ReproIssue15475Remote />);
    expect(baseElement).toBeTruthy();
  });
});
