import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import { TodoList } from './todo-lib';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

describe('TodoList component', () => {
  it('renders without errors', async () => {
    const queryClient = new QueryClient();
    render(
      <QueryClientProvider client={queryClient}>
        <TodoList />
      </QueryClientProvider>
    );
    const buttonElement = screen.getByText('Create Todo');
    expect(buttonElement).toBeInTheDocument();
  });
});
