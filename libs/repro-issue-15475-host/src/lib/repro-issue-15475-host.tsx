import { useState, useEffect, Key } from 'react';
import { useQuery, useMutation } from '@apollo/client';
import { gql } from 'graphql-tag';

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

function TodoList() {
  const [todos, setTodos] = useState<any>([]);

  const { data } = useQuery(GET_TODOS);

  useEffect(() => {
    if (data) {
      setTodos(data.todos);
    }
  }, [data]);

  const [createTodo] = useMutation(CREATE_TODO, {
    onCompleted: (data: any) => {
      setTodos([...todos, data.createTodo]);
    }
  });

  const handleClick = () => {
    createTodo();
  };

  return (
    <div>
      <button onClick={handleClick}>Create Todo</button>
      <ul>
        {todos.map((todo: any, index: Key | null | undefined) => (
          <li key={index}>
            {todo.text} - {todo.done ? 'Done' : 'Not Done'} ({todo.user.name})
          </li>
        ))}
      </ul>
    </div>
  );
}

export default TodoList;
