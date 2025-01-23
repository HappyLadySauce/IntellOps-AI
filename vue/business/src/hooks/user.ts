import { useRouter } from 'vue-router';
import { useUserStore } from '@/store';
export default function useUser() {
  const router = useRouter();
  const userStore = useUserStore();
  const logout = (logoutTo?: string) => {
    userStore.logout();
    const currentRoute = router.currentRoute.value;
    router.push({
      name: logoutTo && typeof logoutTo === 'string' ? logoutTo : 'login',
      query: {
        ...router.currentRoute.value.query,
        redirect: currentRoute.name as string,
      },
    });
  };
  return {
    logout,
  };
}
