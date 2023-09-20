package i18n

const koreanIntroPopupMessage = `
lazygit!를 이용해주셔서 감사합니다. Seriously you rock. Three things to share with you:

 1) lazygit의 기능에 대해 알아보려면 다음 비디오를 참조하세요.
      https://youtu.be/CPLdltN7wgE

 2) 다음 사이트에서 최신 릴리스 노트를 읽어보세요.:
      https://github.com/jesseduffield/lazygit/releases

 3) 만약 당신이 Git을 사용한다면, 그것은 당신을 프로그래머로 만들 것입니다!
	 당신의 도움으로 우리는 lazygit을 더 좋게 만들 수 있습니다, 그러니 기여자가 되는 것을 고려해보세요. 그리고 재미에 참여하세요:
      https://github.com/jesseduffield/lazygit
    또한 오른쪽 하단의 기부 버튼을 클릭하여 저를 후원하고 작업할 내용을 알려주실 수 있습니다.
    또는 저장소에 스타를 눌러 사랑을 공유할 수도 있습니다!
`

// exporting this so we can use it in tests
func koreanTranslationSet() TranslationSet {
	return TranslationSet{
		NotEnoughSpace:                      "패널을 렌더링 할 공간이 부족합니다.",
		DiffTitle:                           "Diff",
		FilesTitle:                          "파일",
		BranchesTitle:                       "브랜치",
		CommitsTitle:                        "커밋",
		StashTitle:                          "Stash",
		UnstagedChanges:                     `Staged되지 않은 변경 내용`,
		StagedChanges:                       `Staged된 변경 내용`,
		MainTitle:                           "메인",
		MergeConfirmTitle:                   "병합",
		StagingTitle:                        "메인 패널 (Staging)",
		MergingTitle:                        "메인 패널 (Merging)",
		NormalTitle:                         "메인 패널 (Normal)",
		LogTitle:                            "로그",
		CommitSummary:                       "커밋 메시지",
		CredentialsUsername:                 "사용자 이름",
		CredentialsPassword:                 "패스워드",
		CredentialsPassphrase:               "SSH키의 passphrase 입력",
		PassUnameWrong:                      "패스워드, passphrase 또는 사용자 이름이 잘못되었습니다.",
		CommitChanges:                       "커밋 변경내용",
		AmendLastCommit:                     "마지맛 커밋 수정",
		AmendLastCommitTitle:                "마지막 커밋 수정",
		SureToAmend:                         "마지막 커밋을 수정하시겠습니까? 그런 다음 커밋 패널에서 커밋 메시지를 변경할 수 있습니다.",
		NoCommitToAmend:                     "Amend 가능한 커밋이 없습니다.",
		CommitChangesWithEditor:             "Git 편집기를 사용하여 변경 내용을 커밋합니다.",
		StatusTitle:                         "상태",
		Menu:                                "메뉴",
		Execute:                             "실행",
		ToggleStaged:                        "Staged 전환",
		ToggleStagedAll:                     "모든 변경을 Staged/unstaged으로 전환",
		ToggleTreeView:                      "파일 트리뷰로 전환",
		OpenMergeTool:                       "Git mergetool를 열기",
		Refresh:                             "새로고침",
		Push:                                "푸시",
		Pull:                                "업데이트",
		Scroll:                              "스크롤",
		MergeConflictsTitle:                 "병합 충돌 내용",
		Checkout:                            "체크아웃",
		FileFilter:                          "파일을 필터하기 (Staged/unstaged)",
		FilterStagedFiles:                   "Staged된 파일만 표시",
		FilterUnstagedFiles:                 "Stage되지 않은 파일만 표시",
		ResetFilter:                         "필터 리셋",
		NoChangedFiles:                      "변경된 파일이 없습니다.",
		SoftReset:                           "소프트 리셋",
		AlreadyCheckedOutBranch:             "브랜치가 이미 체크아웃 되었습니다",
		SureForceCheckout:                   "강제로 체크아웃하시겠습니까? 모든 로컬 변경 사항을 잃게 됩니다.",
		ForceCheckoutBranch:                 "브랜치 강제 체크아웃",
		BranchName:                          "브랜치 이름",
		NewBranchNameBranchOff:              "새 브랜치 이름 (branch is off of '{{.branchName}}')",
		CantDeleteCheckOutBranch:            "체크아웃하는 브랜치는 삭제할 수 없습니다!",
		ForceDeleteBranchMessage:            "'{{.selectedBranchName}}'는 완전히 병합되지 않았습니다. 정말 삭제하시겠습니까?",
		RebaseBranch:                        "체크아웃된 브랜치를 이 브랜치에 리베이스",
		CantRebaseOntoSelf:                  "브랜치를 자기 자신에게 리베이스할 수는 없습니다.",
		CantMergeBranchIntoItself:           "브랜치를 자기 자신에게 병합할 수는 없습니다.",
		ForceCheckout:                       "강제 체크아웃",
		CheckoutByName:                      "이름으로 체크아웃",
		NewBranch:                           "새 브랜치 생성",
		NoBranchesThisRepo:                  "저장소에 브랜치가 존재하지 않습니다.",
		CommitWithoutMessageErr:             "커밋 메시지를 입력하세요.",
		CloseCancel:                         "닫기/취소",
		Confirm:                             "확인",
		Close:                               "닫기",
		Quit:                                "종료",
		SquashDown:                          "Squash down",
		FixupCommit:                         "Fixup commit",
		NoCommitsThisBranch:                 "이 브랜치에 커밋이 없습니다.",
		CannotSquashOrFixupFirstCommit:      "There's no commit below to squash into",
		Fixup:                               "Fixup",
		SureFixupThisCommit:                 "Are you sure you want to 'fixup' this commit? It will be merged into the commit below",
		SureSquashThisCommit:                "Are you sure you want to squash this commit into the commit below?",
		Squash:                              "Squash",
		PickCommit:                          "Pick commit (when mid-rebase)",
		RevertCommit:                        "커밋 되돌리기",
		RewordCommit:                        "커밋메시지 변경",
		DeleteCommit:                        "커밋 삭제",
		MoveDownCommit:                      "커밋을 1개 아래로 이동",
		MoveUpCommit:                        "커밋을 1개 위로 이동",
		EditCommit:                          "커밋을 편집",
		AmendToCommit:                       "Amend commit with staged changes",
		ResetAuthor:                         "Reset commit author",
		SureResetCommitAuthor:               "The author field of this commit will be updated to match the configured user. This also renews the author timestamp. Continue?",
		RenameCommitEditor:                  "에디터에서 커밋메시지 수정",
		Error:                               "오류",
		PickHunk:                            "Pick hunk",
		PickAllHunks:                        "Pick all hunks",
		Undo:                                "되돌리기",
		UndoReflog:                          "되돌리기 (reflog) (실험적)",
		RedoReflog:                          "다시 실행 (reflog) (실험적)",
		Pop:                                 "Pop",
		Drop:                                "Drop",
		Apply:                               "적용",
		NoStashEntries:                      "Stash가 존재하지 않습니다.",
		StashDrop:                           "Stash를 삭제",
		SureDropStashEntry:                  "정말로 Stash를 삭제하시겠습니까?",
		StashPop:                            "Stash를 pop",
		SurePopStashEntry:                   "정말로 Stash를 pop하시겠습니까?",
		StashApply:                          "Stash 적용",
		SureApplyStashEntry:                 "정말로 Stash를 적용하시겠습니까?",
		NoTrackedStagedFilesStash:           "You have no tracked/staged files to stash",
		StashChanges:                        "변경을 Stash",
		RenameStash:                         "Rename stash",
		RenameStashPrompt:                   "Rename stash: {{.stashName}}",
		OpenConfig:                          "설정 파일 열기",
		EditConfig:                          "설정 파일 수정",
		ForcePush:                           "강제 푸시",
		ForcePushPrompt:                     "브랜치가 원격 브랜치에서 분기하고 있습니다. 'esc'를 눌러 취소하거나, 'enter'를 눌러 강제로 푸시하세요.",
		ForcePushDisabled:                   "브랜치가 원격 브랜치에서 분기하고 있습니다. force push가 비활성화 되었습니다.",
		UpdatesRejectedAndForcePushDisabled: "업데이트가 거부되었으며 강제 푸시를 비활성화했습니다.",
		CheckForUpdate:                      "업데이트 확인",
		CheckingForUpdates:                  "업데이트 확인 중...",
		UpdateAvailableTitle:                "새로운 업데이트 사용가능!",
		UpdateAvailable:                     "버전 {{.newVersion}} 을(를) 설치하시겠습니까?",
		UpdateInProgressWaitingStatus:       "업데이트 중",
		UpdateCompletedTitle:                "업데이트 완료!",
		UpdateCompleted:                     "업데이트 설치에 성공했습니다. lazygit를 재시작해주세요.",
		FailedToRetrieveLatestVersionErr:    "버전 정보를 받아오는데 실패했습니다.",
		OnLatestVersionErr:                  "이미 최신 버전을 사용하고 있습니다.",
		MajorVersionErr:                     "새 버전 ({{.newVersion}}) 에 현재 버전({{.currentVersion}}) 과 비교할 때 호환되지 않는 변경 사항이 있습니다.",
		CouldNotFindBinaryErr:               "{{.url}} 에서 바이너리를 찾을 수 없습니다.",
		UpdateFailedErr:                     "업데이트 실패: {{.errMessage}}",
		ConfirmQuitDuringUpdateTitle:        "현재 업데이트 중입니다.",
		ConfirmQuitDuringUpdate:             "현재 업데이트를 진행 중입니다.종료하시겠습니까?",
		MergeToolTitle:                      "병합 도구",
		MergeToolPrompt:                     "정말로 `git mergetool`을 여시겠습니까?",
		IntroPopupMessage:                   koreanIntroPopupMessage,
		GitconfigParseErr:                   `따옴표로 묶이지 않은 '\' 문자가 있어서 Gogit이 gitconfig 파일을 분석하지 못했습니다. 이를 제거하면 문제가 해결됩니다.`,
		EditFile:                            `파일 편집`,
		OpenFile:                            `파일 닫기`,
		IgnoreFile:                          `.gitignore에 추가`,
		RefreshFiles:                        `파일 새로고침`,
		MergeIntoCurrentBranch:              `현재 브랜치에 병합`,
		ConfirmQuit:                         `정말로 종료하시겠습니까?`,
		SwitchRepo:                          `최근에 사용한 저장소로 전환`,
		AllBranchesLogGraph:                 `모든 브랜치 로그 표시`,
		UnsupportedGitService:               `지원되지 않는 Git 서비스입니다.`,
		CreatePullRequest:                   `풀 리퀘스트 생성`,
		CopyPullRequestURL:                  `풀 리퀘스트 URL을 클립보드에 복사`,
		NoBranchOnRemote:                    `브랜치가 원격에 없습니다. 원격에 먼저 푸시해야합니다.`,
		Fetch:                               `Fetch`,
		NoAutomaticGitFetchTitle:            `자동 git 업데이트) 없음`,
		NoAutomaticGitFetchBody:             `Lazygit은 private 저장소에서 "git fetch"를 사용할 수 없습니다. 파일 패널에서 'f'를 사용하여 "git fetch"를 수동으로 실행하세요.`,
		FileEnter:                           `Stage individual hunks/lines for file, or collapse/expand for directory`,
		FileStagingRequirements:             `추적된 파일에 대해 개별 라인만 stage할 수 있습니다.`,
		StageSelection:                      `선택한 행을 staged / unstaged`,
		DiscardSelection:                    `변경을 삭제 (git reset)`,
		ToggleDragSelect:                    `드래그 선택 전환`,
		ToggleSelectHunk:                    `Toggle select hunk`,
		ToggleSelectionForPatch:             `Line(s)을 패치에 추가/삭제`,
		ToggleStagingPanel:                  `패널 전환`,
		ReturnToFilesPanel:                  `파일 목록으로 돌아가기`,
		FastForward:                         `Fast-forward this branch from its upstream`,
		FastForwarding:                      "Fast-forwarding {{.branch}} ...",
		FoundConflictsTitle:                 "Auto-merge failed",
		ViewMergeRebaseOptions:              "View merge/rebase options",
		NotMergingOrRebasing:                "You are currently neither rebasing nor merging",
		RecentRepos:                         "최근에 사용한 저장소",
		MergeOptionsTitle:                   "Merge options",
		RebaseOptionsTitle:                  "Rebase options",
		CommitSummaryTitle:                  "커밋메시지",
		LocalBranchesTitle:                  "브랜치",
		SearchTitle:                         "검색",
		TagsTitle:                           "태그",
		MenuTitle:                           "메뉴",
		RemotesTitle:                        "원격",
		RemoteBranchesTitle:                 "원격 브랜치",
		PatchBuildingTitle:                  "메인 패널 (Patch Building)",
		InformationTitle:                    "정보",
		SecondaryTitle:                      "Secondary",
		ReflogCommitsTitle:                  "Reflog",
		GlobalTitle:                         "글로벌 키 바인딩",
		ConflictsResolved:                   "모든 병합 충돌이 해결되었습니다. 계속 할까요?",
		ConfirmMerge:                        "정말로 '{{.selectedBranch}}' 을(를) '{{.checkedOutBranch}}'에 병합하시겠습니까?",
		FwdNoUpstream:                       "Cannot fast-forward a branch with no upstream",
		FwdNoLocalUpstream:                  "Cannot fast-forward a branch whose remote is not registered locally",
		FwdCommitsToPush:                    "Cannot fast-forward a branch with commits to push",
		ErrorOccurred:                       "오류가 발생했습니다! issue를 작성해 주세요: ",
		NoRoom:                              "Not enough room",
		YouAreHere:                          "현재 위치",
		RewordNotSupported:                  "Rewording commits while interactively rebasing is not currently supported",
		CherryPickCopy:                      "커밋을 복사 (cherry-pick)",
		CherryPickCopyRange:                 "커밋을 범위로 복사 (cherry-pick)",
		PasteCommits:                        "커밋을 붙여넣기 (cherry-pick)",
		SureCherryPick:                      "정말로 복사한 커밋을 이 브랜치에 체리픽하시겠습니까?",
		CherryPick:                          "체리픽",
		Donate:                              "후원",
		AskQuestion:                         "질문하기",
		PrevLine:                            "이전 줄 선택",
		NextLine:                            "다음 줄 선택",
		PrevHunk:                            "이전 hunk를 선택",
		NextHunk:                            "다음 hunk를 선택",
		PrevConflict:                        "이전 충돌을 선택",
		NextConflict:                        "다음 충돌을 선택",
		SelectPrevHunk:                      "이전 hunk를 선택",
		SelectNextHunk:                      "다음 hunk를 선택",
		ScrollDown:                          "아래로 스크롤",
		ScrollUp:                            "위로 스크롤",
		ScrollUpMainPanel:                   "메인 패널을 위로 스크롤",
		ScrollDownMainPanel:                 "메인 패널을 아래로로 스크롤",
		AmendCommitTitle:                    "Amend commit",
		AmendCommitPrompt:                   "Are you sure you want to amend this commit with your staged files?",
		DeleteCommitTitle:                   "커밋 삭제",
		DeleteCommitPrompt:                  "정말로 선택한 커밋을 삭제하시겠습니까?",
		PullingStatus:                       "업데이트 중",
		PushingStatus:                       "푸시 중",
		FetchingStatus:                      "패치 중",
		SquashingStatus:                     "Squashing",
		FixingStatus:                        "Fixing up",
		DeletingStatus:                      "Deleting",
		MovingStatus:                        "Moving",
		RebasingStatus:                      "Rebasing",
		AmendingStatus:                      "Amending",
		CherryPickingStatus:                 "Cherry-picking",
		UndoingStatus:                       "Undoing",
		RedoingStatus:                       "Redoing",
		CheckingOutStatus:                   "Checking out",
		CommittingStatus:                    "Committing",
		CommitFiles:                         "Commit files",
		SubCommitsDynamicTitle:              "커밋 (%s)",
		CommitFilesDynamicTitle:             "Diff files (%s)",
		RemoteBranchesDynamicTitle:          "원격브랜치 (%s)",
		ViewItemFiles:                       "View selected item's files",
		CommitFilesTitle:                    "커밋 파일",
		CheckoutCommitFile:                  "Checkout file",
		DiscardOldFileChange:                "Discard this commit's changes to this file",
		DiscardFileChangesTitle:             "파일 변경 사항 버리기",
		DiscardFileChangesPrompt:            "Are you sure you want to discard this commit's changes to this file? If this file was created in this commit, it will be deleted",
		DisabledForGPG:                      "Feature not available for users using GPG",
		CreateRepo:                          "Git 저장소가 아닙니다. 저장소를 생성하시겠습니까? (y/n): ",
		AutoStashTitle:                      "Autostash?",
		AutoStashPrompt:                     "You must stash and pop your changes to bring them across. Do this automatically? (enter/esc)",
		StashPrefix:                         "Auto-stashing changes for ",
		ViewDiscardOptions:                  "View 'discard changes' options",
		Cancel:                              "취소",
		DiscardAllChanges:                   "모든 변경사항 버리기",
		DiscardUnstagedChanges:              "Discard unstaged changes",
		DiscardAllChangesToAllFiles:         "Nuke working tree",
		DiscardAnyUnstagedChanges:           "Discard unstaged changes",
		DiscardUntrackedFiles:               "Discard untracked files",
		HardReset:                           "Hard reset",
		ViewResetOptions:                    `View reset options`,
		CreateFixupCommitDescription:        `Create fixup commit for this commit`,
		SquashAboveCommits:                  `Squash all 'fixup!' commits above selected commit (autosquash)`,
		SureSquashAboveCommits:              `Are you sure you want to squash all fixup! commits above {{.commit}}?`,
		CreateFixupCommit:                   `Create fixup commit`,
		SureCreateFixupCommit:               `Are you sure you want to create a fixup! commit for commit {{.commit}}?`,
		ExecuteCustomCommand:                "Execute custom command",
		CustomCommand:                       "Custom command:",
		CommitChangesWithoutHook:            "Commit changes without pre-commit hook",
		SkipHookPrefixNotConfigured:         "You have not configured a commit message prefix for skipping hooks. Set `git.skipHookPrefix = 'WIP'` in your config",
		ResetTo:                             `Reset to`,
		PressEnterToReturn:                  "엔터를 눌러 lazygit으로 돌아갑니다.",
		ViewStashOptions:                    "Stash 옵션 보기",
		StashAllChanges:                     "변경사항을 Stash",
		StashStagedChanges:                  "Stash staged changes",
		StashOptions:                        "Stash 옵션",
		NotARepository:                      "Error: must be run inside a git repository",
		Jump:                                "패널로 이동",
		ScrollLeftRight:                     "좌우로 스크롤",
		ScrollLeft:                          "우 스크롤",
		ScrollRight:                         "좌 스크롤",
		DiscardPatch:                        "Patch 버리기",
		DiscardPatchConfirm:                 "You can only build a patch from one commit/stash-entry at a time. Discard current patch?",
		CantPatchWhileRebasingError:         "You cannot build a patch or run patch commands while in a merging or rebasing state",
		ToggleAddToPatch:                    "Toggle file included in patch",
		ToggleAllInPatch:                    "Toggle all files included in patch",
		UpdatingPatch:                       "Updating patch",
		ViewPatchOptions:                    "커스텀 Patch 옵션 보기",
		PatchOptionsTitle:                   "Patch 옵션",
		NoPatchError:                        "No patch created yet. To start building a patch, use 'space' on a commit file or enter to add specific lines",
		EnterFile:                           "Enter file to add selected lines to the patch (or toggle directory collapsed)",
		// ExitCustomPatchBuilder:              ``,
		EnterUpstream:              `'<remote> <branchname>'와 같은 형식으로 입력하세요.`,
		InvalidUpstream:            "Upstream의 형식이 잘못되었습니다.'<remote> <branchname>' 와 같은 형식으로 입력하세요.",
		ReturnToRemotesList:        `원격목록으로 돌아가기`,
		AddNewRemote:               `새로운 Remote 추가`,
		NewRemoteName:              `새로운 Remote 이름:`,
		NewRemoteUrl:               `새로운 Remote URL:`,
		EditRemoteName:             `{{.remoteName}} 의 새로운 Remote 이름 입력:`,
		EditRemoteUrl:              `{{.remoteName}} 의 새로운 Remote URL 입력:`,
		RemoveRemote:               `Remote를 삭제`,
		RemoveRemotePrompt:         "정말로 Remote를 삭제하시겠습니까?",
		DeleteRemoteBranch:         "원격 브랜치를 삭제",
		DeleteRemoteBranchMessage:  "정말로 원격 브랜치를 삭제하시겠습니까?",
		SetUpstream:                "Set as upstream of checked-out branch",
		SetUpstreamTitle:           "Set upstream branch",
		SetUpstreamMessage:         "Are you sure you want to set the upstream branch of '{{.checkedOut}}' to '{{.selected}}'",
		EditRemote:                 "Remote를 수정",
		TagCommit:                  "Tag commit",
		TagMenuTitle:               "태그 작성",
		TagNameTitle:               "태그 이름",
		TagMessageTitle:            "태그 메시지",
		AnnotatedTag:               "Annotated tag",
		LightweightTag:             "Lightweight tag",
		PushTagTitle:               "원격에 태그 '{{.tagName}}' 를 푸시",
		PushTag:                    "태그를 push",
		CreateTag:                  "태그를 생성",
		FetchRemote:                "원격을 업데이트",
		FetchingRemoteStatus:       "원격을 업데이트 중",
		CheckoutCommit:             "커밋을 체크아웃",
		SureCheckoutThisCommit:     "정말로 선택한 커밋을 체크아웃 하시겠습니까?",
		GitFlowOptions:             "Git-flow 옵션 보기",
		NotAGitFlowBranch:          "This does not seem to be a git flow branch",
		NewGitFlowBranchPrompt:     "New {{.branchType}} name:",
		IgnoreTracked:              "Ignore tracked file",
		IgnoreTrackedPrompt:        "Are you sure you want to ignore a tracked file?",
		ViewResetToUpstreamOptions: "View upstream reset options",
		NextScreenMode:             "다음 스크린 모드 (normal/half/fullscreen)",
		PrevScreenMode:             "이전 스크린 모드",
		StartSearch:                "검색 시작",
		Panel:                      "패널",
		Keybindings:                "키 바인딩",
		RenameBranch:               "브랜치 이름 변경",
		NewBranchNamePrompt:        "새로운 브랜치 이름 입력",
		RenameBranchWarning:        "This branch is tracking a remote. This action will only rename the local branch name, not the name of the remote branch. Continue?",
		OpenMenu:                   "매뉴 열기",
		ResetCherryPick:            "Reset cherry-picked (copied) commits selection",
		NextTab:                    "이전 탭",
		PrevTab:                    "다음 탭",
		CantUndoWhileRebasing:      "리베이스중에는 되돌릴 수 없습니다.",
		CantRedoWhileRebasing:      "리베이스중에는 다시 실행할 수 없습니다.",
		MustStashWarning:           "Pulling a patch out into the index requires stashing and unstashing your changes. If something goes wrong, you'll be able to access your files from the stash. Continue?",
		MustStashTitle:             "Must stash",
		ConfirmationTitle:          "확인 패널",
		PrevPage:                   "이전 페이지",
		NextPage:                   "다음 페이지",
		GotoTop:                    "맨 위로 스크롤 ",
		GotoBottom:                 "맨 아래로 스크롤 ",
		FilteringBy:                "Filtering by",
		ResetInParentheses:         "(reset)",
		OpenFilteringMenu:          "View filter-by-path options",
		FilterBy:                   "Filter by",
		ExitFilterMode:             "Stop filtering by path",
		FilterPathOption:           "Enter path to filter by",
		EnterFileName:              "Enter path:",
		FilteringMenuTitle:         "Filtering",
		MustExitFilterModeTitle:    "Command not available",
		MustExitFilterModePrompt:   "Command not available in filtered mode. Exit filtered mode?",
		Diff:                       "Diff",
		EnterRefToDiff:             "Enter ref to diff",
		EnterRefName:               "Ref 입력:",
		ExitDiffMode:               "Diff 모드 종료",
		DiffingMenuTitle:           "Diff",
		SwapDiff:                   "Reverse diff direction",
		OpenDiffingMenu:            "Diff 메뉴 열기",
		// the actual view is the extras view which I intend to give more tabs in future but for now we'll only mention the command log part
		OpenExtrasMenu:                      "명령어 로그 메뉴 열기",
		ShowingGitDiff:                      "Showing output for:",
		CommitDiff:                          "커밋의 iff",
		CopyCommitShaToClipboard:            "커밋 SHA를 클립보드에 복사",
		CommitSha:                           "커밋 SHA",
		CommitURL:                           "커밋 URL",
		CopyCommitMessageToClipboard:        "커밋 메시지를 클립보드에 복사",
		CommitMessage:                       "커밋 메시지",
		CommitAuthor:                        "커밋 작성자",
		CopyCommitAttributeToClipboard:      "커밋 attribute 복사",
		CopyBranchNameToClipboard:           "브랜치명을 클립보드에 복사",
		CopyFileNameToClipboard:             "파일명을 클립보드에 복사",
		CopyCommitFileNameToClipboard:       "커밋한 파일명을 클립보드에 복사",
		CopySelectedTexToClipboard:          "선택한 텍스트를 클립보드에 복사",
		CommitPrefixPatternError:            "Error in commitPrefix pattern",
		NoFilesStagedTitle:                  "파일이 Staged 되지 않았습니다.",
		NoFilesStagedPrompt:                 "파일이 Staged 되지 않았습니다. 모든 파일을 커밋하시겠습니까?",
		BranchNotFoundTitle:                 "브랜치를 찾을 수 없습니다.",
		BranchNotFoundPrompt:                "브랜치를 찾을 수 없습니다. 새로운 브랜치를 생성합니다.",
		DiscardChangeTitle:                  "선택한 라인을 unstaged",
		DiscardChangePrompt:                 "정말로 선택한 라인을 삭제 (git reset) 하시겠습니까? 이 조작은 취소할 수 없습니다.\n이 경고를 비활성화 하려면 설정 파일의 'gui.skipDiscardChangeWarning' 를 true로 설정하세요.",
		CreateNewBranchFromCommit:           "커밋에서 새 브랜치를 만듭니다.",
		BuildingPatch:                       "Building patch",
		ViewCommits:                         "커밋 보기",
		MinGitVersionError:                  "Lazygit 실행을 위해서는 Git 2.20 이후의 버전(2018년 이후의)이 필요합니다. Git를 업데이트 해주세요. 아니면 lazygit이 이전 버전과 더 잘 호환되도록 https://github.com/jesseduffield/lazygit/issues 에 issue를 작성해 주세요.",
		RunningCustomCommandStatus:          "커스텀 명령어 실행",
		SubmoduleStashAndReset:              "Stash uncommitted submodule changes and update",
		AndResetSubmodules:                  "And reset submodules",
		EnterSubmodule:                      "서브모듈 열기",
		CopySubmoduleNameToClipboard:        "서브모듈 이름을 클립보드에 복사",
		RemoveSubmodule:                     "서브모듈 삭제",
		RemoveSubmodulePrompt:               "정말로 서브모듈 '%s'및 해당 디렉토리를 제거하시겠습니까? 이것은 되돌릴 수 없습니다.",
		ResettingSubmoduleStatus:            "서브모듈를 리셋",
		NewSubmoduleName:                    "새로운 서브모듈이름 :",
		NewSubmoduleUrl:                     "새로운 서브모듈의 URL:",
		NewSubmodulePath:                    "새로운 서브모듈의 경로",
		AddSubmodule:                        "새로운 서브모듈 추가",
		AddingSubmoduleStatus:               "새로운 서브모듈 추가",
		UpdateSubmoduleUrl:                  "서브모듈 '%s' 의 URL을 업데이트",
		UpdatingSubmoduleUrlStatus:          "Updating URL",
		EditSubmoduleUrl:                    "서브모듈의 URL을 수정",
		InitializingSubmoduleStatus:         "서브모듈 초기화",
		InitSubmodule:                       "서브모듈 초기화",
		SubmoduleUpdate:                     "서브모듈 업데이트",
		UpdatingSubmoduleStatus:             "서브모듈 업데이트",
		BulkInitSubmodules:                  "서브모듈 일괄 초기화",
		BulkUpdateSubmodules:                "서브모듈 일괄 업데이트",
		BulkDeinitSubmodules:                "Bulk deinit submodules",
		ViewBulkSubmoduleOptions:            "View bulk submodule options",
		BulkSubmoduleOptions:                "Bulk submodule options",
		RunningCommand:                      "Running command",
		SubCommitsTitle:                     "Sub-commits",
		SubmodulesTitle:                     "서브모듈",
		NavigationTitle:                     "List panel navigation",
		SuggestionsCheatsheetTitle:          "추천",
		SuggestionsTitle:                    "추천 (press %s to focus)",
		ExtrasTitle:                         "명령어 로그",
		PushingTagStatus:                    "Pushing tag",
		PullRequestURLCopiedToClipboard:     "풀 리퀘스트의 URL을 클립보드에 복사했습니다.",
		CommitDiffCopiedToClipboard:         "커밋의 Diff를 클립보드에 복사했습니다.",
		CommitSHACopiedToClipboard:          "커밋의 SHA를 클립보드에 복사했습니다.",
		CommitURLCopiedToClipboard:          "커밋의 URL를 클립보드에 복사했습니다.",
		CommitMessageCopiedToClipboard:      "커밋 메시지를 클립보드에 복사했습니다.",
		CommitAuthorCopiedToClipboard:       "커밋 작성자를 클립보드에 복사했습니다.",
		CopiedToClipboard:                   "클립보드에 복사했습니다.",
		ErrCannotEditDirectory:              "디렉토리는 편집할 수 없습니다.",
		ErrStageDirWithInlineMergeConflicts: "병합 충돌이 발생한 파일을 포함하는 디렉토리는 Staged/untaged할 수 없습니다. 병합 충돌을 먼저 해결하세요.",
		ErrRepositoryMovedOrDeleted:         "저장소를 찾을 수 없습니다. 이미 삭제되었거나 이동되었을 가능성이 있습니다. ¯\\_(ツ)_/¯",
		CommandLog:                          "명령어 로그",
		ToggleShowCommandLog:                "명령어 로그 표시 여부 전환",
		FocusCommandLog:                     "명령어 로그에 포커스",
		CommandLogHeader:                    "명령어 로그표시 여부는 '%s' 으로 전환할 수 있습니다.\n",
		RandomTip:                           "랜덤 Tip",
		SelectParentCommitForMerge:          "병합을 위한 상위 커밋 선택",
		ToggleWhitespaceInDiffView:          "공백문자를 Diff 뷰에서 표시 여부 전환",
		IncreaseContextInDiffView:           "Diff 보기의 변경 사항 주위에 표시되는 컨텍스트의 크기를 늘리기",
		DecreaseContextInDiffView:           "Diff 보기의 변경 사항 주위에 표시되는 컨텍스트 크기 줄이기",
		CreatePullRequestOptions:            "풀 리퀘스트 생성 옵션",
		DefaultBranch:                       "기본 브랜치",
		SelectBranch:                        "브랜치를 선택",
		SelectConfigFile:                    "설정파일 선택",
		NoConfigFileFoundErr:                "설정 파일을 찾지 못했습니다.",
		LoadingFileSuggestions:              "파일 제안 로딩 중",
		LoadingCommits:                      "커밋 로딩",
		MustSpecifyOriginError:              "Must specify a remote if specifying a branch",
		GitOutput:                           "Git output:",
		GitCommandFailed:                    "Git command failed. Check command log for details (open with %s)",
		AbortTitle:                          "%s 중지",
		AbortPrompt:                         "정말로 실행중인 %s 를 중지할까요?",
		OpenLogMenu:                         "로그 메뉴 열기",
		LogMenuTitle:                        "커밋 로그 옵션",
		ToggleShowGitGraphAll:               "Toggle show whole git graph (pass the `--all` flag to `git log`)",
		ShowGitGraph:                        "커밋 그래프 표시",
		SortCommits:                         "커밋 정렬",
		CantChangeContextSizeError:          "Cannot change context while in patch building mode because we were too lazy to support it when releasing the feature. If you really want it, please let us know!",
		OpenCommitInBrowser:                 "브라우저에서 커밋 열기",
		ViewBisectOptions:                   "Bisect 옵션 보기",
		ConfirmRevertCommit:                 "Are you sure you want to revert {{.selectedCommit}}?",
		RewordInEditorTitle:                 "커밋 메시지를 에디터에서 수정",
		RewordInEditorPrompt:                "Are you sure you want to reword this commit in your editor?",
		HardResetAutostashPrompt:            "Are you sure you want to hard reset to '%s'? An auto-stash will be performed if necessary.",
		CheckoutPrompt:                      "Are you sure you want to checkout '%s'?",
		UpstreamGone:                        "(upstream gone)",
		Actions: Actions{
			// TODO: combine this with the original keybinding descriptions (those are all in lowercase atm)
			CheckoutCommit:                    "커밋 체크아웃",
			CheckoutTag:                       "태그 체크아웃",
			CheckoutBranch:                    "브랜치 체크아웃",
			ForceCheckoutBranch:               "브랜치 Force 체크아웃",
			Merge:                             "병합",
			RebaseBranch:                      "브랜치 리베이스",
			RenameBranch:                      "브랜치 이름 변경",
			CreateBranch:                      "브랜치 생성",
			CherryPick:                        "(Cherry-pick) 커밋 붙여넣기",
			CheckoutFile:                      "체크아웃 파일",
			DiscardOldFileChange:              "Discard old file change",
			SquashCommitDown:                  "Squash commit down",
			FixupCommit:                       "커밋 Fixup",
			RewordCommit:                      "커밋 Reword",
			DropCommit:                        "커밋 Drop",
			EditCommit:                        "커밋 수정",
			AmendCommit:                       "커밋 Amend",
			ResetCommitAuthor:                 "커밋 작성자 Reset",
			RevertCommit:                      "커밋 Revert",
			CreateFixupCommit:                 "Fixup 커밋 생성",
			SquashAllAboveFixupCommits:        "Squash all above fixup commits",
			CreateLightweightTag:              "Create lightweight tag",
			CreateAnnotatedTag:                "Create annotated tag",
			CopyCommitMessageToClipboard:      "커밋 메시지를 클립보드에 복사",
			CopyCommitDiffToClipboard:         "커밋 diff를 클립보드에 복사",
			CopyCommitSHAToClipboard:          "커밋 SHA를 클립보드에 복사",
			CopyCommitURLToClipboard:          "커밋 URL를 클립보드에 복사",
			CopyCommitAuthorToClipboard:       "커밋 작성자를 클립보드에 복사",
			CopyCommitAttributeToClipboard:    "클립보드에 복사",
			MoveCommitUp:                      "Move commit up",
			MoveCommitDown:                    "Move commit down",
			CustomCommand:                     "Custom command",
			DiscardAllChangesInDirectory:      "Discard all changes in directory",
			DiscardUnstagedChangesInDirectory: "Discard unstaged changes in directory",
			DiscardAllChangesInFile:           "Discard all changes in file",
			DiscardAllUnstagedChangesInFile:   "Discard all unstaged changes in file",
			StageFile:                         "Stage file",
			StageResolvedFiles:                "Stage files whose merge conflicts were resolved",
			UnstageFile:                       "Unstage file",
			UnstageAllFiles:                   "Unstage all files",
			StageAllFiles:                     "Stage all files",
			IgnoreExcludeFile:                 "Ignore file",
			Commit:                            "커밋",
			EditFile:                          "파일 수정",
			Push:                              "푸시",
			Pull:                              "업데이트(Pull)",
			OpenFile:                          "파일 열기",
			StashAllChanges:                   "Stash all changes",
			StashAllChangesKeepIndex:          "Stash all changes and keep index",
			StashStagedChanges:                "Stash staged changes",
			StashUnstagedChanges:              "Stash unstaged changes",
			GitFlowFinish:                     "git flow finish",
			GitFlowStart:                      "git flow start",
			CopyToClipboard:                   "Copy to clipboard",
			CopySelectedTextToClipboard:       "Copy selected text to clipboard",
			RemovePatchFromCommit:             "Remove patch from commit",
			MovePatchToSelectedCommit:         "Move patch to selected commit",
			MovePatchIntoIndex:                "Move patch into index",
			MovePatchIntoNewCommit:            "Move patch into new commit",
			DeleteRemoteBranch:                "Delete remote branch",
			SetBranchUpstream:                 "Set branch upstream",
			AddRemote:                         "Add remote",
			RemoveRemote:                      "Remove remote",
			UpdateRemote:                      "Update remote",
			ApplyPatch:                        "Apply patch",
			Stash:                             "Stash",
			RenameStash:                       "Rename stash",
			RemoveSubmodule:                   "서브모듈 삭제",
			ResetSubmodule:                    "서브모듈 Reset",
			AddSubmodule:                      "서브모듈 추가",
			UpdateSubmoduleUrl:                "서브모듈 URL 업데이트",
			InitialiseSubmodule:               "서브모듈 초기화",
			BulkInitialiseSubmodules:          "Bulk initialise submodules",
			BulkUpdateSubmodules:              "Bulk update submodules",
			BulkDeinitialiseSubmodules:        "Bulk deinitialise submodules",
			UpdateSubmodule:                   "서브모듈 업데이트",
			PushTag:                           "태그 푸시g",
			NukeWorkingTree:                   "Nuke working tree",
			DiscardUnstagedFileChanges:        "Unstaged 파일 변경사항 버리기",
			RemoveUntrackedFiles:              "Untracked 파일 삭제",
			RemoveStagedFiles:                 "Staged 파일 삭제",
			SoftReset:                         "Soft reset",
			MixedReset:                        "Mixed reset",
			HardReset:                         "Hard reset",
			FastForwardBranch:                 "Fast forward branch",
			Undo:                              "되돌리기",
			Redo:                              "다시 실행",
			CopyPullRequestURL:                "풀 리퀘스트 URL 복사",
			OpenMergeTool:                     "병합 도구 열기",
			OpenCommitInBrowser:               "브라우저에서 커밋 열기",
			OpenPullRequest:                   "브라우저에서 풀 리퀘스트 열기",
			StartBisect:                       "Start bisect",
			ResetBisect:                       "Reset bisect",
			BisectSkip:                        "Bisect skip",
			BisectMark:                        "Bisect mark",
		},
		Bisect: Bisect{
			Mark:                        "Mark %s as %s",
			MarkStart:                   "Mark %s as %s (start bisect)",
			SkipCurrent:                 "%s 를 스킵",
			ResetTitle:                  "'git bisect' 를 리셋",
			ResetPrompt:                 "정말로 'git bisect' 를 리셋하시겠습니까?",
			ResetOption:                 "Bisect를 리셋",
			BisectMenuTitle:             "Bisect",
			CompleteTitle:               "Bisect 완료",
			CompletePrompt:              "Bisect complete! The following commit introduced the change:\n\n%s\n\nDo you want to reset 'git bisect' now?",
			CompletePromptIndeterminate: "Bisect complete! Some commits were skipped, so any of the following commits may have introduced the change:\n\n%s\n\nDo you want to reset 'git bisect' now?",
		},
	}
}
