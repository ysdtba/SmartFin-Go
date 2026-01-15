import React, { useState } from 'react';
import { Layout, Menu, Breadcrumb, theme } from 'antd';
import {
  DesktopOutlined,
  PieChartOutlined,
  FileOutlined,
  TeamOutlined,
  UserOutlined,
} from '@ant-design/icons';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';

const { Header, Content, Footer, Sider } = Layout;

type MenuItem = {
  key: React.Key;
  icon?: React.ReactNode;
  children?: MenuItem[];
  label: React.ReactNode;
  path?: string;
};

function getItem(
  label: React.ReactNode,
  key: React.Key,
  icon?: React.ReactNode,
  path?: string,
  children?: MenuItem[],
): MenuItem {
  return {
    key,
    icon,
    children,
    label,
    path,
  } as MenuItem;
}

const items: MenuItem[] = [
  getItem('Dashboard', '1', <PieChartOutlined />, '/'),
  getItem('交易中心', '2', <DesktopOutlined />, '/trade'),
  getItem('资产分析', 'sub1', <UserOutlined />, undefined, [
    getItem('持仓概览', '3', undefined, '/assets'),
    getItem('流水明细', '4', undefined, '/transactions'),
  ]),
  getItem('AI 投研', '9', <FileOutlined />, '/ai-research'),
];

const MainLayout: React.FC = () => {
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();
  const navigate = useNavigate();
  const location = useLocation();

  const handleMenuClick = (e: { key: string }) => {
    // Flatten items to find the clicked one and its path
    // Simplified logic for this demo
    const findPath = (items: MenuItem[], key: string): string | undefined => {
       for (const item of items) {
           if (item.key === key) return item.path;
           if (item.children) {
               const childPath = findPath(item.children, key);
               if (childPath) return childPath;
           }
       }
       return undefined;
    };
    
    const path = findPath(items, e.key);
    if (path) {
        navigate(path);
    }
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <div className="demo-logo-vertical" style={{ height: 32, margin: 16, background: 'rgba(255, 255, 255, 0.2)', borderRadius: 6 }} />
        <Menu 
            theme="dark" 
            defaultSelectedKeys={['1']} 
            mode="inline" 
            items={items} 
            onClick={handleMenuClick}
        />
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer }} />
        <Content style={{ margin: '0 16px' }}>
          <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>SmartFin-Go</Breadcrumb.Item>
            <Breadcrumb.Item>Dashboard</Breadcrumb.Item>
          </Breadcrumb>
          <div
            style={{
              padding: 24,
              minHeight: 360,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}
          >
            <Outlet />
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>
          SmartFin-Go ©{new Date().getFullYear()} Created by FinTech Team
        </Footer>
      </Layout>
    </Layout>
  );
};

export default MainLayout;






