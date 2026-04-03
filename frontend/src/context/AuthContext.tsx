import React, { createContext, useState, useContext, ReactNode } from 'react';

interface AuthContextType {
  userId: string | null;
  isAuthenticated: boolean;
  login: (id: string, accessToken: string, refreshToken: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [userId, setUserId] = useState<string | null>(localStorage.getItem('user_id'));

  const login = (id: string, accessToken: string, refreshToken: string) => {
    localStorage.setItem('user_id', id);
    localStorage.setItem('access_token', accessToken);
    localStorage.setItem('refresh_token', refreshToken);
    setUserId(id);
  };

  const logout = () => {
    localStorage.clear();
    setUserId(null);
  };

  return (
    <AuthContext.Provider value={{ userId, isAuthenticated: !!userId, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) throw new Error('useAuth must be used within an AuthProvider');
  return context;
};
