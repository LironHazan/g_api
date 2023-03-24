import { renderHook } from '@testing-library/react-hooks';
import { DocumentNode, QueryHookOptions, QueryResult, TypedDocumentNode, useQuery } from '@apollo/client';
import { useRepos } from './useRepos';

const useQueryMock = useQuery as jest.MockedFunction<(
  query: DocumentNode | TypedDocumentNode,
  options?: QueryHookOptions<any, Record<string, any>>
) => QueryResult<any, Record<string, any>>>;

jest.mock('@apollo/client');

describe('useRepos', () => {
  const data = {
    user: {
      repositories: {
        nodes: [
          {
            name: 'Repo 1',
            url: 'https://github.com/user/repo1',
            description: 'Description for Repo 1',
          },
          {
            name: 'Repo 2',
            url: 'https://github.com/user/repo2',
            description: 'Description for Repo 2',
          },
        ],
      },
    },
  };

  beforeEach(() => {
    useQueryMock.mockReturnValue({
      loading: false,
      error: null,
      data: data,
    }  as any);
  });

  afterEach(() => {
    jest.resetAllMocks();
  });

  it('should return loading while fetching data', () => {
    useQueryMock.mockReturnValueOnce({
      loading: true,
      error: null,
      data: null,
    }  as any);

    const { result } = renderHook(() => useRepos('user'));
    expect(result.current.loading).toBe(true);
    expect(result.current.error).toBe(null);
    expect(result.current.repos).toEqual([]);
  });

  it('should return an error if the query fails', () => {
    useQueryMock.mockReturnValueOnce({
      loading: false,
      error: new Error('Query failed'),
      data: null,
    }  as any);

    const { result } = renderHook(() => useRepos('user'));
    expect(result.current.loading).toBe(false);
    expect(result.current.error).not.toBe(null);
    expect(result.current.repos).toEqual([]);
  });

  it('should return the list of repositories if the query succeeds', () => {
    const { result } = renderHook(() => useRepos('user'));
    expect(result.current.loading).toBe(false);
    expect(result.current.error).toBe(null);
    expect(result.current.repos).toEqual([
      { id: 'Repo 1', name: 'Repo 1', url: 'https://github.com/user/repo1', description: 'Description for Repo 1' },
      { id: 'Repo 2', name: 'Repo 2', url: 'https://github.com/user/repo2', description: 'Description for Repo 2' },
    ]);
  });
});
