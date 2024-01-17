import Cookies from 'js-cookie'
import { devtools, persist } from 'zustand/middleware'
import { create } from 'zustand'
import { CONFIG } from 'src/config'

export const LOGIN_COOKIE_KEY = CONFIG.LOGIN_COOKIE_KEY

export const UI_SLICE_PERSIST_KEY = 'ui-slice'

interface UIState {
  isLoggingOut: boolean
  setIsLoggingOut: (v: boolean) => void
  accessToken: string
  burgerOpened: boolean
  setBurgerOpened: (v: boolean) => void
}

const useUISlice = create<UIState>()(
  devtools(
    persist(
      (set) => {
        return {
          isLoggingOut: false,
          setIsLoggingOut: (v: boolean) => set((state) => ({ isLoggingOut: v })),
          accessToken: Cookies.get(LOGIN_COOKIE_KEY) ?? '',
          burgerOpened: false,
          setBurgerOpened: (v: boolean) => set((state) => ({ burgerOpened: v })),
        }
      },
      {
        version: 2,
        name: UI_SLICE_PERSIST_KEY,
        partialize(state) {
          const { accessToken, ...rest } = state // always refetch access token from storage
          return rest
        },
      },
    ),
    { enabled: true },
  ),
)

export { useUISlice }

type UIAction = (...args: any[]) => Partial<UIState>
