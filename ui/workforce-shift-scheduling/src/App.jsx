import { useState, useTransition } from 'react';
import { useTranslation } from 'react-i18next';
import ShiftForm from './components/ShiftForm';
import WorkerForm from './components/WorkerForm';
import ScheduleResult from './components/ScheduleResult';
import { getSchedule } from './api';

function App() {
  const { t, i18n } = useTranslation();
  const [shifts, setShifts] = useState([]);
  const [workers, setWorkers] = useState([]);
  const [result, setResult] = useState([]);

  const handleSubmit = async () => {
    try {
      const res = await getSchedule({ shifts, workers });
      setResult(res);
    } catch (err) {
      console.error('API Error:', err);
    }
  };

  return (
    <div className='container mt-4'>
      <div className='d-flex justify-content-between align-items-center mb-3'>
        <h1>{t('title')}</h1>
        <div>
          <button onClick={() => i18n.changeLanguage('it')} className='btn btn-outline-secondary me-2'>IT</button>
          <button onClick={() => i18n.changeLanguage('en')} className='btn btn-outline-secondary'>EN</button>
        </div>
      </div>
      <div className='row'>
        <div className='col-md-6'>
          <ShiftForm onAdd={(s) => setShifts([...shifts, s])} />
        </div>
        <div className='col-md-6'>
          <WorkerForm onAdd={(w) => setWorkers([...workers, w])} />
        </div>
      </div>
      <div className='text-center mt-4'>
        <button className='btn btn-primary' onClick={handleSubmit}>{t('calculate')}</button>
      </div>

      <ScheduleResult result={result} />
    </div>

  );
}

export default App;
