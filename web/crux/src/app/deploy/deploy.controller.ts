import { Body, Controller, UseGuards, UseInterceptors } from '@nestjs/common'
import { concatAll, from, Observable } from 'rxjs'
import { AuditLogLevel } from 'src/decorators/audit-logger.decorators'
import { AuditLoggerInterceptor } from 'src/interceptors/audit-logger.interceptor'
import { GrpcContextLogger } from 'src/interceptors/grpc-context-logger.interceptor'
import { PrismaErrorInterceptor } from 'src/interceptors/prisma-error-interceptor'
import { Empty } from 'src/proto/proto/agent'
import {
  CreateDeploymentRequest,
  CreateEntityResponse,
  CruxDeploymentController,
  CruxDeploymentControllerMethods,
  DeploymentDetailsResponse,
  DeploymentEditEventMessage,
  DeploymentEventListResponse,
  DeploymentListResponse,
  DeploymentProgressMessage,
  IdRequest,
  PatchDeploymentRequest,
  ServiceIdRequest,
  UpdateDeploymentRequest,
  UpdateEntityResponse,
} from 'src/proto/proto/crux'
import { DisableTeamAccessCheck } from 'src/shared/team-access.guard'
import { DeployService } from './deploy.service'
import { DeployCreateTeamAccessGuard } from './guards/deploy.create.team-access.guard'
import { DeployGetByVersionTeamAccessGuard } from './guards/deploy.get-by-version.team-access.guard'
import { DeployTeamAccessGuard } from './guards/deploy.team-access.guard'
import { DeployCreateValidationPipe } from './pipes/deploy.create.pipe'
import { DeleteDeploymentValidationPipe } from './pipes/deploy.delete.pipe'
import { DeployPatchValidationPipe } from './pipes/deploy.patch.pipe'
import { DeployStartValidationPipe } from './pipes/deploy.start.pipe'
import { DeployUpdateValidationPipe } from './pipes/deploy.update.pipe'

@Controller()
@CruxDeploymentControllerMethods()
@UseGuards(DeployTeamAccessGuard)
@UseInterceptors(PrismaErrorInterceptor, GrpcContextLogger, AuditLoggerInterceptor)
export class DeployController implements CruxDeploymentController {
  constructor(private service: DeployService) {}

  @DisableTeamAccessCheck()
  @UseGuards(DeployGetByVersionTeamAccessGuard)
  async getDeploymentsByVersionId(request: IdRequest): Promise<DeploymentListResponse> {
    return await this.service.getDeploymentsByVersionId(request)
  }

  async getDeploymentDetails(request: IdRequest): Promise<DeploymentDetailsResponse> {
    return await this.service.getDeploymentDetails(request)
  }

  async getDeploymentEvents(request: IdRequest): Promise<DeploymentEventListResponse> {
    return await this.service.getDeploymentEvents(request)
  }

  @UseGuards(DeployCreateTeamAccessGuard)
  async createDeployment(
    @Body(DeployCreateValidationPipe) request: CreateDeploymentRequest,
  ): Promise<CreateEntityResponse> {
    return await this.service.createDeployment(request)
  }

  async updateDeployment(
    @Body(DeployUpdateValidationPipe) request: UpdateDeploymentRequest,
  ): Promise<UpdateEntityResponse> {
    return await this.service.updateDeployment(request)
  }

  @AuditLogLevel('no-data')
  async patchDeployment(
    @Body(DeployPatchValidationPipe) request: PatchDeploymentRequest,
  ): Promise<UpdateEntityResponse> {
    return await this.service.patchDeployment(request)
  }

  async deleteDeployment(@Body(DeleteDeploymentValidationPipe) request: IdRequest): Promise<Empty> {
    return await this.service.deleteDeployment(request)
  }

  startDeployment(@Body(DeployStartValidationPipe) request: IdRequest): Observable<DeploymentProgressMessage> {
    return from(this.service.startDeployment(request)).pipe(concatAll())
  }

  @DisableTeamAccessCheck()
  @AuditLogLevel('disabled')
  subscribeToDeploymentEditEvents(request: ServiceIdRequest): Observable<DeploymentEditEventMessage> {
    return this.service.subscribeToDeploymentEditEvents(request)
  }
}
