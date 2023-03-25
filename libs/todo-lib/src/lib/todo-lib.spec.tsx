import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import { rest } from 'msw';
import { setupServer } from 'msw/node';
import TodoList from './todo-lib';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const server = setupServer(
  rest.post('/graphql', (req, res, ctx) => {
    return res(
      ctx.json({
        createTodo: {
          text: 'New Todo',
          done: false,
          user: {
            id: '1',
            name: 'John',
          },
        },
      })
    );
  }),
  rest.get('/graphql', (req, res, ctx) => {
    return res(
      ctx.json({
        todos: [
          {
            text: 'Todo 1',
            done: true,
            user: {
              name: 'John',
            },
          },
          {
            text: 'Todo 2',
            done: false,
            user: {
              name: 'Jane',
            },
          },
        ],
      })
    );
  })
);

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

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

  it('adds a new todo when "Create Todo" button is clicked', async () => {
    const queryClient = new QueryClient();
    render(
      <QueryClientProvider client={queryClient}>
        <TodoList />
      </QueryClientProvider>
    );
    const buttonElement = screen.getByText('Create Todo');
    fireEvent.click(buttonElement);
    await screen.findByText('todo - Not Done (user )');
  });
});
