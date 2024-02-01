import { useState } from 'react';

import { Company } from '@/lib/types';

import Button from '@/ui/Button';
import Modal from '@/ui/Modal';
import CompanyForm from '@/ui/companies/CompanyForm';
import Link from 'next/link';

import style from '@/styles//company/CompanyCard.module.css';

import { updateCompany } from '@/lib/companies';

export default function CompanyCard({ company }: { company: Company }) {
  const [showModal, setShowModal] = useState(false);

  const onClose = () => {
    setShowModal(false);
  }

  return (
    <div className={style.card}>
      <Link
        className={style.link}
        href={`/companies/${company.id}`}
        key={company.id}
      >
        <h2>{company.name}</h2>
        <p>{company.candidate_portal}</p>
      </Link>

      <div className={style.btns}>
        <Button className={style.delete} onClick={() => console.log('delete')}>Delete</Button>

        <Button onClick={() => setShowModal(true)} type="submit">
          Edit
        </Button>
      </div>
      
      {showModal && (
        <Modal onClose={() => setShowModal(false)} title="Update company">
          <CompanyForm action={
            async (formData: FormData) => {
              try {
                // Here company id should never be null
                await updateCompany(company.id!, formData);
                onClose();
                return 'Company updated'
              } catch (error) {
                console.error(error);
                return 'Error updating company'
              }
            }
          } onClose={onClose} company={company} />
        </Modal>
      )}
    </div>
  );
}
