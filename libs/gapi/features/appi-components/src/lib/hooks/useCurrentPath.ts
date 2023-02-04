import { useLocation } from "react-router-dom"

const normalizePathStr = (pathname = ''): string => {
  return pathname.split('/')[1];
}
export const useCurrentPath = () => {
  const location = useLocation();
  return normalizePathStr(location.pathname);
}
