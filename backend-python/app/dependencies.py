from typing import Annotated

from fastapi import Depends
from sqlmodel import Session

from .internal.config.database import get_database_session


DatabaseDep = Annotated[Session, Depends(get_database_session)]
