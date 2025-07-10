import { useState } from 'react';
import { useTranslation } from 'react-i18next';

export default function WorkerForm({ onAdd }) {
  const {t} = useTranslation();
  const [id, setId] = useState('');
  const [name, setName] = useState('');
  const [maxShifts, setMaxShifts] = useState(2);
  const [availability, setAvailability] = useState({});

  const handleAvailability = (e) => {
    const {t} = useTranslation();
    const date = prompt(t('insert_date'));
    const slot = prompt(t('insert_slot'));
    if (date && slot) {
      setAvailability((prev) => ({
        ...prev,
        [date]: [...(prev[date] || []), slot]
      }));
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onAdd({ id, name, max_shifts: parseInt(maxShifts), worked_shifts: 0, availability });
    setId('');
    setName('');
    setAvailability({});
  };

  return (
    <form onSubmit={handleSubmit}>
      <h4>{t('add_worker')}</h4>
      <div className='mb-3'>
        <input placeholder="ID" value={id} onChange={(e) => setId(e.target.value)} required />
      </div>
      <div className='mb-3'>
        <input placeholder="Nome" value={name} onChange={(e) => setName(e.target.value)} required />
      </div>
       <div className='mb-3'>
        <input type="number" value={maxShifts} onChange={(e) => setMaxShifts(e.target.value)} min={1} />
       </div>
      
      <button type="button" onClick={handleAvailability} className='btn btn-outline-secondary'>{t('add_availability')}</button>
      <button type="submit" className='btn btn-outline-secondary'>{t('add_worker')}</button>
    </form>
  );
}
