function useMounted() {
  const isMounted = useRef(false);
  useEffect(() => {
    isMounted.current = true;
    return () => {
      isMounted.current = false;
    }
  }, [isMounted.current]);
  return isMounted.current;
}

export { useMounted }