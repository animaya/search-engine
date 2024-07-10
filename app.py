import streamlit as st
import numpy as np
import pandas as pd
import plotly.graph_objects as go
from streamlit_lottie import st_lottie
import requests

def load_lottieurl(url: str):
    r = requests.get(url)
    if r.status_code != 200:
        return None
    return r.json()

# Настройка страницы
st.set_page_config(page_title="Fit4Success AI", page_icon="🤖", layout="wide")

# Загрузка анимации
lottie_url = "https://assets3.lottiefiles.com/packages/lf20_xGHzgF.json"  # AI-themed animation
lottie_ai = load_lottieurl(lottie_url)

# Пользовательские стили CSS
st.markdown("""
<style>
    @import url('https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;700&display=swap');
    
    html, body, [class*="css"] {
        font-family: 'Roboto', sans-serif;
    }
    
    .main {
        background-color: #111;
        color: #fff;
    }
    
    .stButton>button {
        color: #00ff00;
        background-color: transparent;
        border: 2px solid #00ff00;
        border-radius: 25px;
        padding: 10px 24px;
        font-weight: bold;
        transition: all 0.3s ease;
    }
    
    .stButton>button:hover {
        color: #111;
        background-color: #00ff00;
        box-shadow: 0 0 15px #00ff00;
    }
    
    h1, h2, h3 {
        background: linear-gradient(45deg, #00ffff, #00ff00);
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
    }
    
    .highlight {
        background: linear-gradient(45deg, rgba(0,255,255,0.1), rgba(0,255,0,0.1));
        border: 1px solid rgba(0,255,255,0.2);
        border-radius: 15px;
        padding: 20px;
        margin-bottom: 20px;
    }
    
    .metrics-container {
        background: rgba(0,255,255,0.1);
        border-radius: 15px;
        padding: 20px;
        margin-bottom: 20px;
    }
    
    .metric-value {
        font-size: 2em;
        font-weight: bold;
        color: #00ffff;
    }
    
    .metric-label {
        color: #00ff00;
    }
</style>
""", unsafe_allow_html=True)

# Заголовок
st.title("🤖 Fit4Success: AI-Powered Wellness Revolution")

# Верхняя секция
col1, col2 = st.columns([2, 1])
with col1:
    st.markdown("""
    ## Добро пожаловать в будущее здоровья и фитнеса
    
    Наша AI-программа благополучия использует передовые алгоритмы для достижения:
    - 🧠 Оптимизации физического и ментального потенциала
    - 🔬 Снижения стресса на основе данных
    - 🚀 Квантового скачка в производительности
    
    Присоединяйтесь к технологической революции в сфере здоровья!
    """)
    
    st.button("Активировать AI-ассистента")

with col2:
    st_lottie(lottie_ai, height=300, key="ai")

# Метрики программы
st.header("🔮 AI-метрики программы")
st.markdown('<div class="metrics-container">', unsafe_allow_html=True)
col1, col2, col3, col4 = st.columns(4)
col1.markdown('<p class="metric-value">24/7</p><p class="metric-label">AI Фитнес-коуч</p>', unsafe_allow_html=True)
col2.markdown('<p class="metric-value">99.9%</p><p class="metric-label">Точность анализа питания</p>', unsafe_allow_html=True)
col3.markdown('<p class="metric-value">∞</p><p class="metric-label">Ментальная поддержка</p>', unsafe_allow_html=True)
col4.markdown('<p class="metric-value">1ms</p><p class="metric-label">Скорость обработки данных</p>', unsafe_allow_html=True)
st.markdown('</div>', unsafe_allow_html=True)

# Ключевые преимущества
st.header("💎 Квантовые преимущества")
st.markdown("""
<div class="highlight">
    <h3>🏋️ Нейросетевой анализ тренировок</h3>
    <p>AI анализирует каждое ваше движение и оптимизирует программу в режиме реального времени!</p>
</div>
""", unsafe_allow_html=True)

st.markdown("""
<div class="highlight">
    <h3>🏆 Система геймификации на блокчейне</h3>
    <ul>
        <li>NFT-бейджи за достижения</li>
        <li>Криптовалюта здоровья за выполнение целей</li>
        <li>VR-церемонии награждения в метавселенной Fit4Success</li>
    </ul>
</div>
""", unsafe_allow_html=True)

# График прогресса
st.header("📈 AI-прогнозирование вашего прогресса")

# Генерация данных для графика
dates = pd.date_range(start="2023-06-01", end="2023-12-31", freq="D")
wellness_score = np.cumsum(np.random.randn(len(dates))) + 100  # Random walk starting at 100
wellness_score = np.clip(wellness_score, 50, 150)  # Clip values between 50 and 150

fig = go.Figure()
fig.add_trace(go.Scatter(x=dates, y=wellness_score, mode='lines', name='Wellness Score',
                         line=dict(color='#00ffff', width=2)))

fig.update_layout(
    title='Прогнозируемый AI wellness-score',
    xaxis_title='Дата',
    yaxis_title='Wellness Score',
    paper_bgcolor='rgba(0,0,0,0)',
    plot_bgcolor='rgba(0,0,0,0)',
    font=dict(color='#ffffff'),
    xaxis=dict(showgrid=False),
    yaxis=dict(showgrid=False)
)

st.plotly_chart(fig, use_container_width=True)

# Призыв к действию
st.markdown("""
<div style="background: linear-gradient(45deg, #00ffff, #00ff00); padding: 20px; border-radius: 15px; text-align: center;">
    <h2 style="color: #111;">Готовы к квантовому скачку в здоровье?</h2>
    <p style="color: #111;">Активируйте свой AI-ассистент Fit4Success прямо сейчас!</p>
</div>
""", unsafe_allow_html=True)

if st.button("Инициировать нейронную связь с Fit4Success"):
    st.balloons()
    st.success("Ваш цифровой двойник создан! Ожидайте квантовой телепортации данных.")

# Нижний колонтитул
st.markdown("---")
st.markdown("© 2023 Fit4Success AI. Все права защищены. Работает на квантовых вычислениях.")