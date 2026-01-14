import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import MainLayout from './layouts/MainLayout';
import Dashboard from './pages/Dashboard';
import { ConfigProvider } from 'antd';
import zhCN from 'antd/locale/zh_CN';

const App: React.FC = () => {
  return (
    <ConfigProvider locale={zhCN}>
      <Router>
        <Routes>
          <Route path="/" element={<MainLayout />}>
            <Route index element={<Dashboard />} />
            <Route path="trade" element={<div>交易中心 (Coming Soon)</div>} />
            <Route path="assets" element={<div>持仓概览 (Coming Soon)</div>} />
            <Route path="transactions" element={<div>流水明细 (Coming Soon)</div>} />
            <Route path="ai-research" element={<div>AI 投研 (Coming Soon)</div>} />
          </Route>
        </Routes>
      </Router>
    </ConfigProvider>
  );
};

export default App;
