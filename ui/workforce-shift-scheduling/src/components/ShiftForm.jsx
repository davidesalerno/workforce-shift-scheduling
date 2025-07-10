import { useState } from 'react';
import { useTranslation } from 'react-i18next';

export default function ShiftForm({ onAdd }) {
    const { t } = useTranslation();
    const [date, setDate] = useState('');
    const [slot, setSlot] = useState('morning');
    const [required, setRequired] = useState(1);

    const handleSubmit = (e) => {
        e.preventDefault();
        onAdd({ date, slot, required: parseInt(required) });
        setDate('');
    };

    return (
        <form onSubmit={handleSubmit}>
            <h4>{t('add_shift')}</h4>
            <div className='mb-3'>
                <input type="date" value={date} onChange={(e) => setDate(e.target.value)} required />
            </div>
            <div className='mb-3'>
                <select value={slot} onChange={(e) => setSlot(e.target.value)}>
                    <option value="morning">{t('morning')}</option>
                    <option value="evening">{t('evening')}</option>
                </select>
            </div>
            <div className='mb-3'>
                <input type="number" value={required} onChange={(e) => setRequired(e.target.value)} min={1} required />

            </div>

            <button type="submit" className='btn btn-outline-secondary'>{t('submit_shift')}</button>
        </form>
    );
}
