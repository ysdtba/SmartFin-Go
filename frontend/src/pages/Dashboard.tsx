import React from 'react';
import { Card, Col, Row, Statistic } from 'antd';
import { ArrowUpOutlined, ArrowDownOutlined } from '@ant-design/icons';

const Dashboard: React.FC = () => (
  <div>
    <h2>市场概览</h2>
    <Row gutter={16}>
      <Col span={8}>
        <Card bordered={false}>
          <Statistic
            title="总资产 (USD)"
            value={112893.00}
            precision={2}
            valueStyle={{ color: '#3f8600' }}
            prefix={<ArrowUpOutlined />}
            suffix="%"
          />
        </Card>
      </Col>
      <Col span={8}>
        <Card bordered={false}>
          <Statistic
            title="今日收益"
            value={9.3}
            precision={2}
            valueStyle={{ color: '#cf1322' }}
            prefix={<ArrowDownOutlined />}
            suffix="%"
          />
        </Card>
      </Col>
      <Col span={8}>
        <Card bordered={false}>
          <Statistic
            title="AI 分析信号"
            value={5}
            suffix="/ 10 (积极)"
          />
        </Card>
      </Col>
    </Row>
    <div style={{ marginTop: 24 }}>
        <h3>系统公告</h3>
        <p>欢迎使用 SmartFin-Go 智能投研系统。</p>
    </div>
  </div>
);

export default Dashboard;





