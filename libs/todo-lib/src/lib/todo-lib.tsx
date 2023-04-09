import { useState, Key } from 'react';
import { gql } from 'graphql-tag';
import graphqlRequestClient from './graphqlRequestClient';
import { useQuery, useMutation } from '@tanstack/react-query';
import { Todo } from './gql/graphql';

const GET_TODOS = gql`
  query findTodos {
    todos {
      text
      done
      user {
        name
      }
    }
  }
`;

const CREATE_TODO = gql`
  mutation createTodo {
    createTodo(input: { text: "todo", userId: "1" }) {
      user {
        id
        name
      }
      text
      done
    }
  }
`;

interface TodosResponse {
  todos: Todo[];
}

export function TodoList() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useQuery({ queryKey: ['posts'], queryFn: async () => {
    const { todos }: TodosResponse = await graphqlRequestClient.request(GET_TODOS)
    setTodos(todos)
    return todos
  } })

  const { mutate }  = useMutation<any>({
    mutationFn: () => graphqlRequestClient.request(CREATE_TODO),
    onSuccess: ({ createTodo }) => { setTodos([...todos, createTodo]) },
  });

  const handleClick = () => {
    mutate();
  };

  return (
    <div>
      <button onClick={handleClick}>Create Todo</button>
      <ul>
        {todos?.map((todo: Todo, index: Key | null | undefined) => (
          <li key={index}>
            {todo.text} - {todo.done ? 'Done' : 'Not Done'} ({todo.user.name})
          </li>
        ))}
      </ul>
    </div>
  );
}
