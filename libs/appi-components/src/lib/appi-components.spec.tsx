import { render } from '@testing-library/react';

import AppiComponents from './appi-components';

describe('AppiComponents', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<AppiComponents />);
    expect(baseElement).toBeTruthy();
  });
});
