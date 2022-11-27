import styled from '@emotion/styled';

/* eslint-disable-next-line */
export interface AppiComponentsProps {}

const StyledAppiComponents = styled.div`
  color: pink;
`;

export function AppiComponents(props: AppiComponentsProps) {
  return (
    <StyledAppiComponents>
      <h1>Welcome to AppiComponents!</h1>
    </StyledAppiComponents>
  );
}

export default AppiComponents;
